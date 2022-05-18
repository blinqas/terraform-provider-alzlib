package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/matt-FFFFFF/alzlib"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"directory": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("ALZLIB_DIR", ""),
					Description: "The directory containing the ALZ lib files.",
				},
			},

			DataSourcesMap: map[string]*schema.Resource{
				"alzlib_archetypes": dataSourceArchetypes(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

// type apiClient struct {
// 	// Add whatever fields, client or connection info, etc. here
// 	// you would need to setup to communicate with the upstream
// 	// API.
// }

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		directory := d.Get("directory").(string)
		az, err := alzlib.NewAlzLib(directory)
		if err != nil {
			diag.Errorf("Failed to read supplied directory: %s. %s", directory, err)
		}
		return az, nil
	}
}
