package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccArchetypesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccArchetypesDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.alzlib_archetypes.test", "archetypes.es_root.name", "es_root"),
					resource.TestCheckResourceAttr("data.alzlib_archetypes.test", "archetypes.es_root.policy_definitions.%", "104"),
				),
			},
		},
	})
}

const testAccArchetypesDataSourceConfig = `
data "alzlib_archetypes" "test" {}
`
