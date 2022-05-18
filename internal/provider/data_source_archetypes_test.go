package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceArchetype(t *testing.T) {
	//t.Skip("data source not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceArchetype,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.alzlib_archetype.test", "output", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccDataSourceArchetype = `
data "alzlib_archetypes" "test" {}
`
