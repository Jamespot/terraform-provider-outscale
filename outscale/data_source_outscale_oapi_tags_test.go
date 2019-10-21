package outscale

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccOutscaleOAPITagsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			skipIfNoOAPI(t)
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOAPITagsDataSourceConfig(),
			},
		},
	})
}

// Lookup based on InstanceID
func testAccOAPITagsDataSourceConfig() string {
	return `
		data "outscale_tags" "web" {
			filter {
				name   = "resource_type"
				values = ["instance"]
			}
		}
	`
}
