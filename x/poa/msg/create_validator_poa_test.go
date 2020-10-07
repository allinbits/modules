package msg

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	pk1      = ed25519.GenPrivKey().PubKey()
	addr1    = pk1.Address()
	valAddr1 = sdk.ValAddress(addr1)

	emptyAddr   sdk.ValAddress
	emptyPubkey crypto.PubKey
)

func TestMsgCreateValidatorPOA(t *testing.T) {
	tests := []struct {
		name, moniker, identity, website, securityContact, details string
		validatorAddr                                              sdk.ValAddress
		pubkey                                                     crypto.PubKey
		expectPass                                                 bool
	}{
		{"basic good", "a", "b", "c", "d", "e", valAddr1, pk1, true},
		{"empty name", "", "", "c", "", "", valAddr1, pk1, false},
		{"empty description", "", "", "", "", "", valAddr1, pk1, false},
		{"empty address", "a", "b", "c", "d", "e", emptyAddr, pk1, false},
		{"empty pubkey", "a", "b", "c", "d", "e", valAddr1, emptyPubkey, true},
	}

	for _, tc := range tests {
		description := stakingtypes.NewDescription(tc.moniker, tc.identity, tc.website, tc.securityContact, tc.details)
		msg := NewMsgCreateValidatorPOA(tc.moniker, tc.validatorAddr, tc.pubkey, description, sdk.AccAddress(addr1))
		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", tc.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", tc.name)
		}
	}
}
