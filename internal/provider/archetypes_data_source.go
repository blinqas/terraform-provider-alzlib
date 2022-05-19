package provider

import (
	"context"
	"log"

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
		MarkdownDescription: "Example data source",

		Attributes: map[string]tfsdk.Attribute{
			"data": {
				Type: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": types.StringType,
						// "policy_definitions": types.ObjectType{
						// 	AttrTypes: map[string]attr.Type{
						// 		"name":         types.StringType,
						// 		"display_name": types.StringType,
						// 	},
						// },
					},
				},
			},
		},
	}, nil
}

func (t archetypesDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return archetypesDataSource{
		provider: provider,
	}, diags
}

type archetypesDataSourceData struct {
	Data archetypesDataSourceDataData `tfsdk:"data"`
}

type archetypesDataSourceDataData struct {
	Name types.String `tfsdk:"name"`
	//PolicyDefinitions archetypesDataSourceDataDataPolicyDefinitions `tfsdk:"policy_definitions"`
}

type archetypesDataSourceDataDataPolicyDefinitions struct {
	Name        types.String `tfsdk:"name"`
	DisplayName types.String `tfsdk:"display_name"`
}

type archetypesDataSource struct {
	provider provider
}

func (d archetypesDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var data archetypesDataSourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	log.Printf("got here")

	// d.provider.client contains the client that was created by the provider server

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("got here")

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// example, err := d.provider.client.ReadExample(...)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// ddpd := archetypesDataSourceDataDataPolicyDefinitions{
	// 	Name:        types.String{Value: "example"},
	// 	DisplayName: types.String{Value: "Example"},
	// }

	dd := archetypesDataSourceDataData{
		Name: types.String{Value: "namevalue"},
		//PolicyDefinitions: ddpd,
	}
	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	dataValue := make(map[string]archetypesDataSourceDataData)
	dataValue["test"] = dd
	data.Data = dd

	//diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
