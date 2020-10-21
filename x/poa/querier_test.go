package poa

import (
	"fmt"
	"testing"

	"github.com/allinbits/modules/x/poa/keeper"
	"github.com/allinbits/modules/x/poa/msg"
	"github.com/allinbits/modules/x/poa/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestQueryValidatorPOA(t *testing.T) {
	ctx, k := keeper.MakeTestCtxAndKeeper(t)
	var cdc = codec.New()

	name := "name"
	valPubKey := keeper.MakeTestPubKey(keeper.SamplePubKey)
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())

	validator := types.NewValidator(
		name,
		valAddr,
		valPubKey,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
	)

	k.SetValidator(ctx, name, validator)

	bz, _ := cdc.MarshalJSON(types.NewQueryValidatorParams(name))

	query := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/validator-poa/%s", types.QuerierRoute, name),
		Data: bz,
	}

	_, err := queryValidator(ctx, query, k)
	require.NoError(t, err)
}

func TestQueryVote(t *testing.T) {
	ctx, k := keeper.MakeTestCtxAndKeeper(t)
	var cdc = codec.New()

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

	msg := msg.NewMsgVoteValidator(name, valAddr, true, accAddr)
	handleMsgVoteValidator(ctx, msg, k)

	bz, _ := cdc.MarshalJSON(types.NewQueryVoteParams(name, valAddr.String()))

	query := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/vote-poa/%s", types.QuerierRoute, name),
		Data: bz,
	}

	_, err := queryVote(ctx, query, k)
	require.NoError(t, err)
}
