package outscale

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	oscgo "github.com/outscale/osc-sdk-go/v2"
	"github.com/terraform-providers/terraform-provider-outscale/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccOutscaleOAPIVirtualRoutePropagation_basic(t *testing.T) {
	t.Parallel()
	rBgpAsn := acctest.RandIntRange(64512, 65534)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOAPIVirtualRoutePropagationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleOAPIVpnRoutePropagationConfig(rBgpAsn),
				Check: resource.ComposeTestCheckFunc(
					testAccOutscaleOAPIVpnRoutePropagation(
						"outscale_virtual_gateway_route_propagation.outscale_virtual_gateway_route_propagation",
					),
				),
			},
		},
	})
}

func testAccCheckOAPIVirtualRoutePropagationDestroy(s *terraform.State) error {
	oscapi := testAccProvider.Meta().(*OutscaleClient).OSCAPI

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "outscale_virtual_gateway_route_propagation" {
			continue
		}

		var resp oscgo.ReadVirtualGatewaysResponse
		var err error

		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			resp, _, err = oscapi.VirtualGatewayApi.ReadVirtualGateways(context.Background()).ReadVirtualGatewaysRequest(oscgo.ReadVirtualGatewaysRequest{
				Filters: &oscgo.FiltersVirtualGateway{VirtualGatewayIds: &[]string{rs.Primary.Attributes["gateway_id"]}},
			}).Execute()
			if err != nil {
				return utils.CheckThrottling(err)
			}
			return nil
		})

		if err != nil {
			return err
		}

		if len(resp.GetVirtualGateways()) > 0 {
			err = resource.Retry(5*time.Minute, func() *resource.RetryError {
				_, _, err := oscapi.VirtualGatewayApi.DeleteVirtualGateway(context.Background()).DeleteVirtualGatewayRequest(oscgo.DeleteVirtualGatewayRequest{
					VirtualGatewayId: resp.GetVirtualGateways()[0].GetVirtualGatewayId(),
				}).Execute()
				if err != nil {
					if strings.Contains(err.Error(), utils.ResourceNotFound) {
						return resource.RetryableError(err)
					}
					return utils.CheckThrottling(err)
				}
				return nil
			})

			if err != nil {
				return fmt.Errorf("ERROR => %s", err)
			}

		} else {
			return nil
		}
	}
	return nil
}

func testAccOutscaleOAPIVpnRoutePropagation(routeProp string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[routeProp]
		if !ok {
			return fmt.Errorf("Not found: %s", routeProp)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		return nil
	}
}

func testAccOutscaleOAPIVpnRoutePropagationConfig(rBgpAsn int) string {
	return fmt.Sprintf(`
		resource "outscale_virtual_gateway" "outscale_virtual_gateway" {
 connection_type = "ipsec.1"
}

resource "outscale_net" "outscale_net" {
    ip_range = "10.0.0.0/18"
}
resource "outscale_virtual_gateway_link" "outscale_virtual_gateway_link" {
    virtual_gateway_id = outscale_virtual_gateway.outscale_virtual_gateway.virtual_gateway_id
    net_id              = outscale_net.outscale_net.net_id
}

resource "outscale_route_table" "outscale_route_table" {
    net_id = outscale_net.outscale_net.net_id
    tags {
     key = "name"
     value = "terraform-RT"
    }
}

resource "outscale_virtual_gateway_route_propagation" "outscale_virtual_gateway_route_propagation" {
virtual_gateway_id = outscale_virtual_gateway_link.outscale_virtual_gateway_link.virtual_gateway_id
    route_table_id  = outscale_route_table.outscale_route_table.route_table_id
    enable = true 
}	
	`)
}
