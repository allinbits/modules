package poa

import (
	"testing"

	"github.com/allinbits/modules/x/poa/keeper"
	"github.com/allinbits/modules/x/poa/msg"
	"github.com/allinbits/modules/x/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

func TestHandleCreateValidatorPOA(t *testing.T) {
	ctx, k := keeper.MakeTestCtxAndKeeper(t)

	name := "name"
	valPubKey := keeper.MakeTestPubKey(keeper.SamplePubKey)
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())
	accAddr := sdk.AccAddress(valPubKey.Address().Bytes())

	msg := msg.NewMsgCreateValidatorPOA(
		name,
		valAddr,
		valPubKey,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
		accAddr,
	)

	res, err := handleMsgCreateValidatorPOA(ctx, msg, k)

	// Validate msg was handled correctly but checking the size of the store
	allVals := k.GetAllValidators(ctx)
	require.Equal(t, 1, len(allVals))

	// No errors and res is populates
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestHandleVoteValidatorPOA(t *testing.T) {
	ctx, k := keeper.MakeTestCtxAndKeeper(t)

	name := "name"
	valPubKey := keeper.MakeTestPubKey(keeper.SamplePubKey)
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())
	accAddr := sdk.AccAddress(valPubKey.Address().Bytes())

	validator := types.NewValidator(
		name,
		valAddr,
		valPubKey,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
	)

	k.SetValidator(ctx, name, validator)
	k.CalculateValidatorVotes(ctx)
	k.ApplyAndReturnValidatorSetUpdates(ctx)

	msg := msg.NewMsgVoteValidator("name", valAddr, true, accAddr)

	res, err := handleMsgVoteValidator(ctx, msg, k)

	// Validate msg was handled correctly but checking the size of the store
	allVotes := k.GetAllVotes(ctx)
	require.Equal(t, 1, len(allVotes))

	// No errors and res is populated
	require.NoError(t, err)
	require.NotNil(t, res)
}
