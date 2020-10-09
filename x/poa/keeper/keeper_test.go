package keeper

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/allinbits/modules/x/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func TestKeeperGenericFunctions(t *testing.T) {
	ctx, keeper := MakeTestCtxAndKeeper(t)

	// Create test validator
	valPubKey := MakeTestPubKey(SamplePubKey)
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())

	validator := types.NewValidator(
		"name",
		valAddr,
		valPubKey,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
	)

	// TODO: split into multiple test cases

	// Set a value in the store
	keeper.Set(ctx, []byte(validator.Name), types.ValidatorsKey, validator)

	// Check the store to see if the item was saved
	_, found := keeper.Get(ctx, []byte(validator.Name), types.ValidatorsKey, keeper.UnmarshalValidator)
	require.True(t, found)

	// Get all items from the store
	allVals := keeper.GetAll(ctx, types.ValidatorsKey, keeper.UnmarshalValidator)
	require.Equal(t, 1, len(allVals))
}
