package keeper

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/allinbits/modules/x/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func TestKeeperUpdateValidatorSetFunctions(t *testing.T) {
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

	valPubKey2 := MakeTestPubKey(SamplePubKey2)
	valAddr2 := sdk.ValAddress(valPubKey.Address().Bytes())

	validator2 := types.NewValidator(
		"name2",
		valAddr2,
		valPubKey2,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
	)

	// Set a value in the store
	keeper.SetValidator(ctx, validator.Name, validator)
	keeper.CalculateValidatorVotes(ctx)

	updates := keeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.Equal(t, 1, len(updates))

	keeper.SetValidator(ctx, validator2.Name, validator2)

	vote := types.NewVote(
		valAddr,
		"name2",
		true,
	)

	vote2 := types.NewVote(
		valAddr2,
		"name2",
		true,
	)

	// Validator 1 votes for validator 2 to join the consensus
	keeper.SetVote(ctx, vote)

	// Validator 2 votes for validator 2 to join the consensus
	keeper.SetVote(ctx, vote2)
	keeper.CalculateValidatorVotes(ctx)

	// Validator 2 joins the consensus
	updates = keeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.Equal(t, 1, len(updates))

	// No updates to the set
	updates = keeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.Equal(t, 0, len(updates))
}

func TestKeeperCalculateValidatorVoteFunction(t *testing.T) {
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

	valPubKey2 := MakeTestPubKey(SamplePubKey2)
	valAddr2 := sdk.ValAddress(valPubKey.Address().Bytes())

	validator2 := types.NewValidator(
		"name2",
		valAddr2,
		valPubKey2,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
	)

	// Set a value in the store
	keeper.SetValidator(ctx, validator.Name, validator)
	keeper.ApplyAndReturnValidatorSetUpdates(ctx)
	keeper.CalculateValidatorVotes(ctx)

	// Check to see if validator is accepted
	retVal, found := keeper.GetValidator(ctx, validator.Name)
	require.True(t, retVal.Accepted)
	require.True(t, found)

	// Set the second validator and assert its not accepted
	keeper.SetValidator(ctx, validator2.Name, validator2)
	retVal, found = keeper.GetValidator(ctx, validator2.Name)
	require.False(t, retVal.Accepted)

	vote := types.NewVote(
		valAddr,
		"name2",
		true,
	)

	// Validator 1 votes for validator 2 to join the consensus
	keeper.SetVote(ctx, vote)
	keeper.CalculateValidatorVotes(ctx)

	// Validator 2's accepted value is set to false
	retVal, found = keeper.GetValidator(ctx, validator2.Name)
	require.True(t, retVal.Accepted)
}
