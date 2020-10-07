package keeper

import (
	"github.com/allinbits/modules/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultParamspace = types.ModuleName
)

// GetParams returns the total set of poa parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the poa parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}
