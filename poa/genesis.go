package poa

import (
	"github.com/allinbits/modules/poa/keeper"
	"github.com/allinbits/modules/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) []abci.ValidatorUpdate {
	// set the params
	k.SetParams(ctx, data.Params)

	for _, validator := range data.Validators {
		k.SetValidator(ctx, validator.Name, validator)
	}

	return k.ApplyAndReturnValidatorSetUpdates(ctx)
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	// TODO: Define logic for exporting state
	return types.DefaultGenesisState()
}
