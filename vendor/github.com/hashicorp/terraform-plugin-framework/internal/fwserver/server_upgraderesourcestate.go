package fwserver

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// UpgradeResourceStateRequest is the framework server request for the
// UpgradeResourceState RPC.
type UpgradeResourceStateRequest struct {
	// TODO: Create framework defined type that is not protocol specific.
	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/340
	RawState *tfprotov6.RawState

	ResourceSchema tfsdk.Schema
	ResourceType   tfsdk.ResourceType
	Version        int64
}

// UpgradeResourceStateResponse is the framework server response for the
// UpgradeResourceState RPC.
type UpgradeResourceStateResponse struct {
	Diagnostics   diag.Diagnostics
	UpgradedState *tfsdk.State
}

// UpgradeResourceState implements the framework server UpgradeResourceState RPC.
func (s *Server) UpgradeResourceState(ctx context.Context, req *UpgradeResourceStateRequest, resp *UpgradeResourceStateResponse) {
	if req == nil {
		return
	}

	// No UpgradedState to return. This could return an error diagnostic about
	// the odd scenario, but seems best to allow Terraform CLI to handle the
	// situation itself in case it might be expected behavior.
	if req.RawState == nil {
		return
	}

	// Terraform CLI can call UpgradeResourceState even if the stored state
	// version matches the current schema. Presumably this is to account for
	// the previous terraform-plugin-sdk implementation, which handled some
	// state fixups on behalf of Terraform CLI. When this happens, we do not
	// want to return errors for a missing ResourceWithUpgradeState
	// implementation or an undefined version within an existing
	// ResourceWithUpgradeState implementation as that would be confusing
	// detail for provider developers. Instead, the framework will attempt to
	// roundtrip the prior RawState to a State matching the current Schema.
	//
	// TODO: To prevent provider developers from accidentially implementing
	// ResourceWithUpgradeState with a version matching the current schema
	// version which would never get called, the framework can introduce a
	// unit test helper.
	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/113
	if req.Version == req.ResourceSchema.Version {
		logging.FrameworkTrace(ctx, "UpgradeResourceState request version matches current Schema version, using framework defined passthrough implementation")

		resourceSchemaType := req.ResourceSchema.TerraformType(ctx)

		rawStateValue, err := req.RawState.Unmarshal(resourceSchemaType)

		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Previously Saved State for UpgradeResourceState",
				"There was an error reading the saved resource state using the current resource schema.\n\n"+
					"If this resource state was last refreshed with Terraform CLI 0.11 and earlier, it must be refreshed or applied with an older provider version first. "+
					"If you manually modified the resource state, you will need to manually modify it to match the current resource schema. "+
					"Otherwise, please report this to the provider developer:\n\n"+err.Error(),
			)
			return
		}

		resp.UpgradedState = &tfsdk.State{
			Schema: req.ResourceSchema,
			Raw:    rawStateValue,
		}

		return
	}

	// Always instantiate new Resource instances.
	logging.FrameworkDebug(ctx, "Calling provider defined ResourceType NewResource")
	resource, diags := req.ResourceType.NewResource(ctx, s.Provider)
	logging.FrameworkDebug(ctx, "Called provider defined ResourceType NewResource")

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceWithUpgradeState, ok := resource.(tfsdk.ResourceWithUpgradeState)

	if !ok {
		resp.Diagnostics.AddError(
			"Unable to Upgrade Resource State",
			"This resource was implemented without an UpgradeState() method, "+
				fmt.Sprintf("however Terraform was expecting an implementation for version %d upgrade.\n\n", req.Version)+
				"This is always an issue with the Terraform Provider and should be reported to the provider developer.",
		)
		return
	}

	logging.FrameworkTrace(ctx, "Resource implements ResourceWithUpgradeState")

	logging.FrameworkDebug(ctx, "Calling provider defined Resource UpgradeState")
	resourceStateUpgraders := resourceWithUpgradeState.UpgradeState(ctx)
	logging.FrameworkDebug(ctx, "Called provider defined Resource UpgradeState")

	// Panic prevention
	if resourceStateUpgraders == nil {
		resourceStateUpgraders = make(map[int64]tfsdk.ResourceStateUpgrader, 0)
	}

	resourceStateUpgrader, ok := resourceStateUpgraders[req.Version]

	if !ok {
		resp.Diagnostics.AddError(
			"Unable to Upgrade Resource State",
			"This resource was implemented with an UpgradeState() method, "+
				fmt.Sprintf("however Terraform was expecting an implementation for version %d upgrade.\n\n", req.Version)+
				"This is always an issue with the Terraform Provider and should be reported to the provider developer.",
		)
		return
	}

	upgradeResourceStateRequest := tfsdk.UpgradeResourceStateRequest{
		RawState: req.RawState,
	}

	if resourceStateUpgrader.PriorSchema != nil {
		logging.FrameworkTrace(ctx, "Initializing populated UpgradeResourceStateRequest state from provider defined prior schema and request RawState")

		priorSchemaType := resourceStateUpgrader.PriorSchema.TerraformType(ctx)

		rawStateValue, err := req.RawState.Unmarshal(priorSchemaType)

		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Previously Saved State for UpgradeResourceState",
				fmt.Sprintf("There was an error reading the saved resource state using the prior resource schema defined for version %d upgrade.\n\n", req.Version)+
					"Please report this to the provider developer:\n\n"+err.Error(),
			)
			return
		}

		upgradeResourceStateRequest.State = &tfsdk.State{
			Raw:    rawStateValue,
			Schema: *resourceStateUpgrader.PriorSchema,
		}
	}

	upgradeResourceStateResponse := tfsdk.UpgradeResourceStateResponse{
		State: tfsdk.State{
			Schema: req.ResourceSchema,
			// Raw is intentionally not set.
		},
	}

	// To simplify provider logic, this could perform a best effort attempt
	// to populate the response State by looping through all Attribute/Block
	// by calling the equivalent of SetAttribute(GetAttribute()) and skipping
	// any errors.

	logging.FrameworkDebug(ctx, "Calling provider defined StateUpgrader")
	resourceStateUpgrader.StateUpgrader(ctx, upgradeResourceStateRequest, &upgradeResourceStateResponse)
	logging.FrameworkDebug(ctx, "Called provider defined StateUpgrader")

	resp.Diagnostics.Append(upgradeResourceStateResponse.Diagnostics...)

	if resp.Diagnostics.HasError() {
		return
	}

	if upgradeResourceStateResponse.DynamicValue != nil {
		logging.FrameworkTrace(ctx, "UpgradeResourceStateResponse DynamicValue set, overriding State")

		upgradedStateValue, err := upgradeResourceStateResponse.DynamicValue.Unmarshal(req.ResourceSchema.TerraformType(ctx))

		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Upgrade Resource State",
				fmt.Sprintf("After attempting a resource state upgrade to version %d, the provider returned state data that was not compatible with the current schema.\n\n", req.Version)+
					"This is always an issue with the Terraform Provider and should be reported to the provider developer:\n\n"+err.Error(),
			)
			return
		}

		resp.UpgradedState = &tfsdk.State{
			Schema: req.ResourceSchema,
			Raw:    upgradedStateValue,
		}

		return
	}

	if upgradeResourceStateResponse.State.Raw.Type() == nil || upgradeResourceStateResponse.State.Raw.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Upgraded Resource State",
			fmt.Sprintf("After attempting a resource state upgrade to version %d, the provider did not return any state data. ", req.Version)+
				"Preventing the unexpected loss of resource state data. "+
				"This is always an issue with the Terraform Provider and should be reported to the provider developer.",
		)
		return
	}

	resp.UpgradedState = &upgradeResourceStateResponse.State
}
