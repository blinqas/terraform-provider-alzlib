package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ tfsdk.DataSourceType = archetypesDataSourceType{}
var _ tfsdk.DataSource = archetypesDataSource{}

type archetypesDataSourceType struct{}

func (t archetypesDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Archetypes data from the provider",

		Attributes: map[string]tfsdk.Attribute{
			// The 'id' attribute is needed for acceptance testing
			"id": {
				Type:     types.Int64Type,
				Computed: true,
			},
			"archetypes": {
				Computed: true,
				Type: types.MapType{
					ElemType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"name":               types.StringType,
							"policy_definitions": policyDefinitionType(),
						},
					},
				},
			},
		},
	}, nil
}

func policyDefinitionType() types.MapType {
	return types.MapType{
		ElemType: types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":         types.StringType,
				"display_name": types.StringType,
				"policy_type":  types.StringType,
				"mode":         types.StringType,
				"description":  types.StringType,
				"policy_rule":  types.StringType,
				"metadata":     types.StringType,
				"parameters":   types.StringType,
			},
		},
	}
}

func (t archetypesDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return archetypesDataSource{
		provider: provider,
	}, diags
}

type archetypesDataSource struct {
	provider provider
}

type archetypesDataSourceData struct {
	Id         int64                    `tfsdk:"id"`
	Archetypes map[string]archetypeData `tfsdk:"archetypes"`
}

func (d archetypesDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// example, err := d.provider.client.ReadExample(...)

	// if err != nil {
	// 	resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	// 	return
	// }

	// Create the data structure that will be stored in the state
	// We need an Id field to run the acceptance tests.
	// Since there can only be one of these data sources per provider instance,
	// we can fix this as a constant.
	data := archetypesDataSourceData{
		Id: 0,
	}

	archs := make(map[string]archetypeData)

	for ak := range d.provider.client.Archetypes {
		archs[ak] = archetypeData{
			Name:              types.String{Value: ak},
			PolicyDefinitions: map[string]policyDefinitionsData{},
		}

		for pdk, pdv := range d.provider.client.Archetypes[ak].PolicyDefinitions {
			pdd := policyDefinitionsData{
				Name:        types.String{Value: pdk},
				DisplayName: types.String{Value: *pdv.Properties.DisplayName},
				PolicyType:  types.String{Value: *pdv.Type},
				Mode:        types.String{Value: *pdv.Properties.Mode},
				Description: types.String{Value: *pdv.Properties.Description},
			}

			policyRule := pdv.Properties.PolicyRule.(map[string]interface{})
			policyRuleStr := flattenJSON(policyRule)
			if policyRuleStr == "" {
				resp.Diagnostics.AddError(fmt.Sprintf("Error generating archetype %s", ak), fmt.Sprintf("Unable to read policy rule in policy definition %s", pdk))
			}
			pdd.PolicyRule = types.String{Value: policyRuleStr}

			pdd.Metadata = types.String{Value: flattenJSON(pdv.Properties.Metadata)}

			parametersStr, err := flattenParameterDefinitionsValueToString(pdv.Properties.Parameters)
			if err != nil {
				resp.Diagnostics.AddError(fmt.Sprintf("Error generating archetype %s", ak), fmt.Sprintf("Unable to read policy parameters in policy definition %s", pdk))
			}
			pdd.Parameters = types.String{Value: parametersStr}

			archs[ak].PolicyDefinitions[pdk] = pdd
		}
	}

	data.Archetypes = archs
	diags := resp.State.Set(ctx, &data)

	resp.Diagnostics.Append(diags...)
}

func flattenParameterDefinitionsValueToString(input map[string]*armpolicy.ParameterDefinitionsValue) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	result, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	compactJson := bytes.Buffer{}
	if err := json.Compact(&compactJson, result); err != nil {
		return "", err
	}

	return compactJson.String(), nil
}
