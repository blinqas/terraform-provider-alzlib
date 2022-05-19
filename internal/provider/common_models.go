package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type archetypeData struct {
	Name              types.String                     `tfsdk:"name"`
	PolicyDefinitions map[string]policyDefinitionsData `tfsdk:"policy_definitions"`
}

type policyDefinitionsData struct {
	Name        types.String `tfsdk:"name"`
	DisplayName types.String `tfsdk:"display_name"`
}
