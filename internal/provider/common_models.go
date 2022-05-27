package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type archetypeData struct {
	Name              types.String                     `tfsdk:"name"`
	PolicyDefinitions map[string]policyDefinitionsData `tfsdk:"policy_definitions"`
}

type policyDefinitionsData struct {
	Name        types.String `tfsdk:"name"`
	DisplayName types.String `tfsdk:"display_name"`
	PolicyType  types.String `tfsdk:"policy_type"`
	Mode        types.String `tfsdk:"mode"`
	Description types.String `tfsdk:"description"`
	PolicyRule  types.String `tfsdk:"policy_rule"`
	Metadata    types.String `tfsdk:"metadata"`
	Parameters  types.String `tfsdk:"parameters"`
}
