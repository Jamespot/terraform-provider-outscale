package outscale

import (
	"fmt"
	"log"
	"time"

	"github.com/terraform-providers/terraform-provider-outscale/osc/oapi"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

// Creates a network interface in the specified subnet
func dataSourceOutscaleOAPINics() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceOutscaleOAPINicsRead,
		Schema: getDSOAPINicsSchema(),
	}
}

func getDSOAPINicsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		//  This is attribute part for schema Nic
		"filter": dataSourceFiltersSchema(),
		"nics": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"account_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"description": &schema.Schema{
						Type:     schema.TypeString,
						Computed: true,
					},
					"is_source_dest_checked": {
						Type:     schema.TypeBool,
						Computed: true,
					},
					"link_nic": {
						Type:     schema.TypeMap,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"delete_on_vm_deletion": {
									Type:     schema.TypeBool,
									Computed: true,
								},
								"device_number": {
									Type:     schema.TypeInt,
									Computed: true,
								},
								"nic_link_id": {
									Type:     schema.TypeString,
									Computed: true,
								},
								"state": {
									Type:     schema.TypeString,
									Computed: true,
								},
								"vm_account_id": {
									Type:     schema.TypeString,
									Computed: true,
								},
								"vm_id": {
									Type:     schema.TypeString,
									Computed: true,
								},
							},
						},
					},
					"link_public_ip": {
						Type:     schema.TypeSet,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"link_public_ip_id": {
									Type:     schema.TypeString,
									Computed: true,
								},
								"public_dns_name": {
									Type:     schema.TypeString,
									Computed: true,
								},
								"public_ip": {
									Type:     schema.TypeString,
									Computed: true,
								},
								"public_ip_account_id": {
									Type:     schema.TypeString,
									Computed: true,
								},
								"public_ip_id": {
									Type:     schema.TypeString,
									Computed: true,
								},
							},
						},
					},
					"mac_address": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"net_id": &schema.Schema{
						Type:     schema.TypeString,
						Computed: true,
					},
					"nic_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"private_dns_name": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"private_ips": {
						Type:     schema.TypeSet,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"is_primary": {
									Type:     schema.TypeBool,
									Computed: true,
								},
								"link_public_ip": {
									Type:     schema.TypeSet,
									Computed: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"link_public_ip_id": {
												Type:     schema.TypeString,
												Computed: true,
											},
											"public_dns_name": {
												Type:     schema.TypeString,
												Computed: true,
											},
											"public_ip": {
												Type:     schema.TypeString,
												Computed: true,
											},
											"public_ip_account_id": {
												Type:     schema.TypeString,
												Computed: true,
											},
											"public_ip_id": {
												Type:     schema.TypeString,
												Computed: true,
											},
										},
									},
								},
								"private_dns_name": {
									Type:     schema.TypeString,
									Computed: true,
								},
								"private_ip": {
									Type:     schema.TypeString,
									Computed: true,
								},
							},
						},
					},
					"security_groups": {
						Type:     schema.TypeList,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"security_group_id": {
									Type:     schema.TypeString,
									Computed: true,
								},
								"security_group_name": {
									Type:     schema.TypeString,
									Computed: true,
								},
							},
						},
					},
					"security_group_ids": &schema.Schema{
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					"state": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"subnet_id": &schema.Schema{
						Type:     schema.TypeString,
						Computed: true,
					},
					"sub_region_name": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"tags": &schema.Schema{
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"key": {
									Type:     schema.TypeString,
									Optional: true,
									Computed: true,
								},
								"value": {
									Type:     schema.TypeString,
									Optional: true,
									Computed: true,
								},
							},
						},
					}},
			},
		},
		"request_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

//Read Nic
func dataSourceOutscaleOAPINicsRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OAPI

	filters, filtersOk := d.GetOk("filter")
	if filtersOk == false {
		return fmt.Errorf("filters, or owner must be assigned, or nic_id must be provided")
	}

	params := oapi.ReadNicsRequest{}
	if filtersOk {
		params.Filters = buildOutscaleOAPIDataSourceNicFilters(filters.(*schema.Set))
	}

	var resp *oapi.POST_ReadNicsResponses
	var err error

	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		resp, err = conn.POST_ReadNics(params)
		return resource.RetryableError(err)
	})

	if err != nil {
		return fmt.Errorf("Error reading Network Interface Cards : %s", err)
	}

	if resp.OK.Nics == nil {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again")
	}

	if len(resp.OK.Nics) == 0 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again")
	}

	nics := resp.OK.Nics

	return resourceDataAttrSetter(d, func(set AttributeSetter) error {
		d.SetId(resource.UniqueId())

		if err := set("nics", getOAPIVMNetworkInterfaceSet(nics)); err != nil {
			log.Printf("[DEBUG] NICS ERR %+v", err)
			return err
		}

		return d.Set("request_id", resp.OK.ResponseContext.RequestId)
	})
}
