package poa

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/allinbits/modules/poa/keeper"
	"github.com/allinbits/modules/poa/msg"
	"github.com/allinbits/modules/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

// TODO: Should be able to import SamplePubKey & MakeTestPubKey from keeper
// TODO: for some reason Go is complaining here? Any ideas?

// ----------------------------------------------------

const (
	SamplePubKey = "b7a3c12dc0c8c748ab07525b701122b88bd78f600c76342d27f25e5f92444cde"
)

func MakeTestPubKey(pk string) (res crypto.PubKey) {
	var buffer bytes.Buffer
	buffer.WriteString(pk)

	pkBytes, err := hex.DecodeString(buffer.String())
	if err != nil {
		panic(err)
	}
	var pkEd ed25519.PubKeyEd25519
	copy(pkEd[:], pkBytes)
	return pkEd
}

// ----------------------------------------------------

func TestHandleCreateValidatorPOA(t *testing.T) {
	ctx, keeper := keeper.MakeTestCtxAndKeeper(t)

	name := "name"
	valPubKey := MakeTestPubKey(SamplePubKey)
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())
	accAddr := sdk.AccAddress(valPubKey.Address().Bytes())

	msg := msg.NewMsgCreateValidatorPOA(
		name,
		valAddr,
		valPubKey,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
		accAddr,
	)

	res, err := handleMsgCreateValidatorPOA(ctx, msg, keeper)

	// Validate msg was handled correctly but checking the size of the store
	allVals := keeper.GetAllValidators(ctx)
	require.Equal(t, 1, len(allVals))

	// No errors and res is populates
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestHandleVoteValidatorPOA(t *testing.T) {
	ctx, keeper := keeper.MakeTestCtxAndKeeper(t)

	name := "name"
	valPubKey := MakeTestPubKey(SamplePubKey)
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())
	accAddr := sdk.AccAddress(valPubKey.Address().Bytes())

	validator := types.NewValidator(
		name,
		valAddr,
		valPubKey,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
	)

	keeper.SetValidator(ctx, name, validator)
	keeper.CalculateValidatorVotes(ctx)
	keeper.ApplyAndReturnValidatorSetUpdates(ctx)

	msg := msg.NewMsgVoteValidator("name", valAddr, true, accAddr)

	res, err := handleMsgVoteValidator(ctx, msg, keeper)

	// Validate msg was handled correctly but checking the size of the store
	allVotes := keeper.GetAllVotes(ctx)
	require.Equal(t, 1, len(allVotes))

	// No errors and res is populated
	require.NoError(t, err)
	require.NotNil(t, res)
}
