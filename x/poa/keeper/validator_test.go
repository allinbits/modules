package keeper

import (
	"testing"

	"github.com/allinbits/modules/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeeperValidatorFunctions(t *testing.T) {
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

	// Set a validator in the store
	keeper.SetValidator(ctx, validator.Name, validator)

	// Check the store to see if the validator was saved
	retVal, found := keeper.GetValidator(ctx, validator.Name)
	assert.Equal(t, validator, retVal, "Should return the correct validator from the store")
	require.True(t, found)

	// Get all validators from the store
	allVals := keeper.GetAllValidators(ctx)
	require.Equal(t, 1, len(allVals))
}
