package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceArchetypes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceArchetypesRead,
		Schema: map[string]*schema.Schema{
			"output": schemaPolicyDefinitions(),
		},
	}
}

func dataSourceArchetypesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	//az := m.(*alzlib.AlzLib)
	return diags
}

func schemaArchetypeMap() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeMap,
		Computed:    true,
		Description: "The collection of archetypes, returned as a map",
		Elem: &schema.Schema{
			Type: schema.TypeSet,
			Type: schema.TypeSet,
		},
	}
}

func schemaPolicyDefinitions() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeMap,
		Computed:    true,
		Description: "The collection of policy definitions in the archetype, returned as a map by policy definition name",
		Elem: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"policy_type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"mode": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"policy_rule": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"metadata": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}
