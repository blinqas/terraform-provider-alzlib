package toproto5

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/internal/fwserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

// UpgradeResourceStateResponse returns the *tfprotov5.UpgradeResourceStateResponse
// equivalent of a *fwserver.UpgradeResourceStateResponse.
func UpgradeResourceStateResponse(ctx context.Context, fw *fwserver.UpgradeResourceStateResponse) *tfprotov5.UpgradeResourceStateResponse {
	if fw == nil {
		return nil
	}

	proto5 := &tfprotov5.UpgradeResourceStateResponse{
		Diagnostics: Diagnostics(fw.Diagnostics),
	}

	upgradedState, diags := State(ctx, fw.UpgradedState)

	proto5.Diagnostics = append(proto5.Diagnostics, Diagnostics(diags)...)
	proto5.UpgradedState = upgradedState

	return proto5
}
