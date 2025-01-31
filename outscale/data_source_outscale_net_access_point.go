package outscale

import (
	"context"
	"fmt"
	"time"

	oscgo "github.com/outscale/osc-sdk-go/v2"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-outscale/utils"
)

func napdSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"net_access_point_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"net_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"service_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"state": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tag_key": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"tag_value": {
			Type:     schema.TypeString,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tags": dataSourceTagsSchema(),
		"route_table_ids": {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"request_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func dataSourceOutscaleNetAccessPoint() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOutscaleNetAccessPointRead,

		Schema: getDataSourceSchemas(napdSchema()),
	}
}

func dataSourceOutscaleNetAccessPointRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	napid, napidOk := d.GetOk("net_access_point_ids")
	filters, filtersOk := d.GetOk("filter")
	filter := new(oscgo.FiltersNetAccessPoint)

	if !napidOk && !filtersOk {
		return fmt.Errorf("One of filters, or net_access_point_ids must be assigned")
	}

	if filtersOk {
		filter = buildOutscaleDataSourcesNAPFilters(filters.(*schema.Set))
	} else {
		filter = &oscgo.FiltersNetAccessPoint{
			NetAccessPointIds: &[]string{napid.(string)},
		}
	}

	req := &oscgo.ReadNetAccessPointsRequest{
		Filters: filter,
	}

	var resp oscgo.ReadNetAccessPointsResponse
	var err error

	err = resource.Retry(30*time.Second, func() *resource.RetryError {
		resp, _, err = conn.NetAccessPointApi.ReadNetAccessPoints(
			context.Background()).
			ReadNetAccessPointsRequest(*req).Execute()
		if err != nil {
			return utils.CheckThrottling(err)
		}
		return nil
	})

	if err != nil {
		return err
	}

	if err := utils.IsResponseEmptyOrMutiple(len(resp.GetNetAccessPoints()), "NetAccessPoint"); err != nil {
		return err
	}

	nap := resp.GetNetAccessPoints()[0]

	d.Set("net_access_point_id", nap.NetAccessPointId)
	d.Set("route_table_ids", flattenStringList(nap.RouteTableIds))
	d.Set("net_id", nap.NetId)
	d.Set("service_name", nap.ServiceName)
	d.Set("state", nap.State)
	d.Set("tags", tagsOSCAPIToMap(nap.GetTags()))

	id := *nap.NetAccessPointId
	d.SetId(id)

	return nil
}
