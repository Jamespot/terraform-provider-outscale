package outscale

import (
	"context"
	"fmt"
	"testing"
	"time"

	oscgo "github.com/outscale/osc-sdk-go/v2"
	"github.com/terraform-providers/terraform-provider-outscale/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccOutscaleOAPIVpnGatewayAttachment_basic(t *testing.T) {
	var vpc oscgo.Net
	var vgw oscgo.VirtualGateway
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: "outscale_virtual_gateway_link.test",
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOAPIVpnGatewayAttachmentConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPILinExists(
						"outscale_net.test",
						&vpc),
					testAccCheckOAPIVirtualGatewayExists(
						"outscale_virtual_gateway.test",
						&vgw),
				),
			},
		},
	})
}

func TestAccResourceVpnGatewayAttachment_importBasic(t *testing.T) {
	resourceName := "outscale_virtual_gateway_link.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOAPIVpnGatewayAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOAPIVpnGatewayAttachmentConfig,
			},
			{
				ResourceName:            resourceName,
				ImportStateIdFunc:       testAccCheckVpnGatewayAttachmentImportStateIDFunc(resourceName),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request_id"},
			},
		},
	})
}

func testAccCheckVpnGatewayAttachmentImportStateIDFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return rs.Primary.ID, nil
	}
}

func TestAccAWSOAPIVpnGatewayAttachment_deleted(t *testing.T) {
	t.Parallel()
	var vgw oscgo.VirtualGateway
	var vpc oscgo.Net

	testDeleted := func(n string) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			_, ok := s.RootModule().Resources[n]
			if ok {
				return fmt.Errorf("expected vpn gateway attachment resource %q to be deleted", n)
			}
			return nil
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: "outscale_virtual_gateway_link.test",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckOAPIVpnGatewayAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOAPIVpnGatewayAttachmentConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPILinExists(
						"outscale_net.test",
						&vpc),
					testAccCheckOAPIVirtualGatewayExists(
						"outscale_virtual_gateway.test",
						&vgw),
					testAccCheckOAPIVpnGatewayAttachmentExists(
						"outscale_virtual_gateway_link.test",
						&vgw),
				),
			},
			{
				Config: testAccNoOAPIVpnGatewayAttachmentConfig,
				Check: resource.ComposeTestCheckFunc(
					testDeleted("outscale_virtual_gateway_link.test"),
				),
			},
		},
	})
}

func testAccCheckOAPIVpnGatewayAttachmentExists(n string, vgw *oscgo.VirtualGateway) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		vpcID := rs.Primary.Attributes["net_id"]
		vgwID := rs.Primary.Attributes["virtual_gateway_id"]

		if len(vgw.GetNetToVirtualGatewayLinks()) == 0 {
			return fmt.Errorf("vpn gateway %q has no attachments", vgwID)
		}

		if vgw.GetNetToVirtualGatewayLinks()[0].GetState() != "attached" {
			return fmt.Errorf("Expected VPN Gateway %q to be in attached state, but got: %q",
				vgwID, vgw.GetNetToVirtualGatewayLinks()[0].GetState())
		}

		if vgw.GetNetToVirtualGatewayLinks()[0].GetNetId() != vpcID {
			return fmt.Errorf("Expected VPN Gateway %q to be attached to VPC %q, but got: %q",
				vgwID, vpcID, vgw.GetNetToVirtualGatewayLinks()[0].GetNetId())
		}

		return nil
	}
}

func testAccCheckOAPIVpnGatewayAttachmentDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*OutscaleClient).OSCAPI

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "outscale_virtual_gateway_link" {
			continue
		}

		vgwID := rs.Primary.Attributes["vpn_gateway_id"]

		var resp oscgo.ReadVirtualGatewaysResponse
		var err error

		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			resp, _, err = conn.VirtualGatewayApi.ReadVirtualGateways(context.Background()).ReadVirtualGatewaysRequest(oscgo.ReadVirtualGatewaysRequest{
				Filters: &oscgo.FiltersVirtualGateway{VirtualGatewayIds: &[]string{vgwID}},
			}).Execute()
			if err != nil {
				return utils.CheckThrottling(err)
			}
			return nil
		})

		if err != nil {
			return err
		}
		vgw := resp.GetVirtualGateways()[0]
		if len(vgw.GetNetToVirtualGatewayLinks()) > 0 {
			if vgw.GetNetToVirtualGatewayLinks()[0].GetState() != "detached" {
				return fmt.Errorf("Expected VPN Gateway %q to be in detached state, but got: %q",
					vgwID, vgw.GetNetToVirtualGatewayLinks()[0].GetState())
			}
		}
	}

	return nil
}

const testAccNoOAPIVpnGatewayAttachmentConfig = `
	resource "outscale_net" "test" {
		ip_range = "10.0.0.0/16"
	}

	resource "outscale_virtual_gateway" "test" {
		connection_type = "ipsec.1"
}
`

const testAccOAPIVpnGatewayAttachmentConfig = `
resource "outscale_virtual_gateway" "test" {
 connection_type = "ipsec.1"
}
resource "outscale_net" "test" {
    ip_range = "10.0.0.0/18"
}
resource "outscale_virtual_gateway_link" "test" {
    virtual_gateway_id = outscale_virtual_gateway.test.virtual_gateway_id
    net_id              = outscale_net.test.net_id
}
`
