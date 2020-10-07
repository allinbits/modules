package msg

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// We need to register this msg in the types package to satisfy the types.msg interface

// Vote messages
type MsgVoteValidator struct {
	Name    string         `json:"name"`
	Voter   sdk.ValAddress `json:"voter"`
	InFavor bool           `json:"infavor"`
	Owner   sdk.AccAddress `json:"owner"`
}

func NewMsgVoteValidator(name string, voter sdk.ValAddress, inFavor bool, owner sdk.AccAddress) MsgVoteValidator {
	return MsgVoteValidator{
		Name:    name,
		Voter:   voter,
		InFavor: inFavor,
		Owner:   owner,
	}
}

// Route should return the name of the module
func (msg MsgVoteValidator) Route() string { return "poa" }

// Type should return the action
func (msg MsgVoteValidator) Type() string { return "vote_validator_poa" }

// ValidateBasic runs stateless checks on the message
func (msg MsgVoteValidator) ValidateBasic() error {
	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty name")
	}
	if msg.Voter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "voter not specified")
	}
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgVoteValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// GetSignBytes encodes the message for signing
func (msg MsgVoteValidator) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgVoteValidator{}, "poa/MsgVoteValidator", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
