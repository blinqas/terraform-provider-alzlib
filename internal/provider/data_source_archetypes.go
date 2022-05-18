package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceArchetypes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceArchetypeRead,
		Schema: map[string]*schema.Schema{
			"output": {
				Type:        schema.TypeMap,
				Description: "The collection of archetypes, returned as a map",
				Computed:    true,
				Elem: map[string]*schema.Schema{
					"policy_definitions": schemaPolicyDefinitions(),
				},
			},
		},
	}
}

func dataSourceArchetypeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	//az := m.(*alzlib.AlzLib)
	return diags
}

func schemaPolicyDefinitions() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeMap,
		Computed:    true,
		Description: "The collection of policy definitions in teh archetype, returned as a map by policy definition name",
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
				ValidateFunc: validation.StringInSlice([]string{
					"All",
					"Indexed",
					"Microsoft.ContainerService.Data",
					"Microsoft.CustomerLockbox.Data",
					"Microsoft.DataCatalog.Data",
					"Microsoft.KeyVault.Data",
					"Microsoft.Kubernetes.Data",
					"Microsoft.MachineLearningServices.Data",
					"Microsoft.Network.Data",
					"Microsoft.Synapse.Data",
				}, true),
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
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
			},

			"metadata": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
			},
		},
	}
}
