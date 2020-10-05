package msg

import (
	"github.com/stretchr/testify/require"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// test ValidateBasic for MsgDelegate
func TestMsgVoteValidator(t *testing.T) {
	tests := []struct {
		name       string
		candidate  string
		voter      sdk.ValAddress
		inFavor    bool
		expectPass bool
	}{
		{"basic good", "asd", valAddr1, true, true},
		{"infavor", "asd", valAddr1, true, true},
		{"empty name", "", valAddr1, true, false},
		{"empty voter", "asd", emptyAddr, true, false},
	}

	for _, tc := range tests {
		msg := NewMsgVoteValidator(tc.candidate, tc.voter, tc.inFavor, sdk.AccAddress(addr1))
		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", tc.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", tc.name)
		}
	}
}
