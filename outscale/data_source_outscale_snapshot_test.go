package outscale

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccOutscaleOAPISnapshotDataSource_basic(t *testing.T) {
	t.Parallel()
	region := os.Getenv("OUTSCALE_REGION")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOutscaleOAPISnapshotDataSourceConfig(region),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPISnapshotDataSourceID("data.outscale_snapshot.snapshot"),
					resource.TestCheckResourceAttr("data.outscale_snapshot.snapshot", "volume_size", "1"),
				),
			},
		},
	})
}

func TestAccOutscaleOAPISnapshotDataSource_multipleFilters(t *testing.T) {
	t.Parallel()
	region := os.Getenv("OUTSCALE_REGION")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOutscaleOAPISnapshotDataSourceConfigWithMultipleFilters(region),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPISnapshotDataSourceID("data.outscale_snapshot.snapshot"),
					resource.TestCheckResourceAttr("data.outscale_snapshot.snapshot", "volume_size", "10"),
				),
			},
		},
	})
}

func testAccCheckOutscaleOAPISnapshotDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find snapshot data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Snapshot data source ID not set")
		}
		return nil
	}
}

func testAccCheckOutscaleOAPISnapshotDataSourceConfig(region string) string {
	return fmt.Sprintf(`
		resource "outscale_volume" "example" {
			subregion_name = "%sa"
			size           = 1
		}

		resource "outscale_snapshot" "snapshot" {
			volume_id = "${outscale_volume.example.id}"
		}

		data "outscale_snapshot" "snapshot" {
			snapshot_id = "${outscale_snapshot.snapshot.id}"
		}
	`, region)
}

func testAccCheckOutscaleOAPISnapshotDataSourceConfigWithMultipleFilters(region string) string {
	return fmt.Sprintf(`
		resource "outscale_volume" "external1" {
			subregion_name = "%sa"
			size           = 10
		}

		resource "outscale_snapshot" "snapshot" {
			volume_id = "${outscale_volume.external1.id}"
		}

		data "outscale_snapshot" "snapshot" {
			snapshot_id = "${outscale_snapshot.snapshot.id}"

			filter {
				name   = "volume_sizes"
				values = ["10"]
			}
		}
	`, region)
}
