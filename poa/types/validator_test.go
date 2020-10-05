package types

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	SamplePubKey  = "b7a3c12dc0c8c748ab07525b701122b88bd78f600c76342d27f25e5f92444cde"
	SamplePubKey2 = "38f49d1fdd929524fe3be04fbfbc98c01ce5b94891d9a9031c3b78fdad8ff039"
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

func TestValidatorEqual(t *testing.T) {
	valPubKey := MakeTestPubKey(SamplePubKey)
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())

	val1 := NewValidator(
		"name",
		valAddr,
		valPubKey,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
	)

	val2 := NewValidator(
		"name",
		valAddr,
		valPubKey,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
	)

	// Test two validator are equal when created with the same data
	require.True(t, val1.Address.Equals(val2.Address))

	valPubKey3 := MakeTestPubKey(SamplePubKey2)
	valAddr3 := sdk.ValAddress(valPubKey3.Address().Bytes())
	val3 := NewValidator(
		"name",
		valAddr3,
		valPubKey3,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
	)

	// Test two validators are no equal when create with different data
	require.False(t, val1.Address.Equals(val3.Address))
}

func TestABCIValidatorUpdate(t *testing.T) {
	valPubKey := MakeTestPubKey(SamplePubKey)
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())

	val1 := NewValidator(
		"name",
		valAddr,
		valPubKey,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
	)

	abciVal := val1.ABCIValidatorUpdate(10)

	// Test the validator updates correctly when being pased via ABCI
	require.Equal(t, tmtypes.TM2PB.PubKey(val1.PubKey), abciVal.PubKey)
	require.Equal(t, int64(10), abciVal.Power)
}
