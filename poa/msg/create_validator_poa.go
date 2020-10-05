package msg

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/crypto"
)

// We need to register this msg in the types package to satisfy the types.msg interface: see types/msg.go

// MsgCreateValidatorPOA defines the message that creates a consensus validator
type MsgCreateValidatorPOA struct {
	Name        string                   `json:"name"`
	Address     sdk.ValAddress           `json:"address"`
	PubKey      crypto.PubKey            `json:"pubkey"`
	Description stakingtypes.Description `json:"description"`
	Owner       sdk.AccAddress           `json:"owner"`
}

// NOTE: can we use same type and different marshal fn
// msgCreateValidator is used in the custom unmarshal functon to unmarshal the public key
type msgCreateValidatorJSON struct {
	Name        string                   `json:"name"`
	Address     sdk.ValAddress           `json:"address"`
	PubKey      string                   `json:"pubkey"`
	Description stakingtypes.Description `json:"description"`
	Owner       sdk.AccAddress           `json:"owner"`
}

func NewMsgCreateValidatorPOA(name string, address sdk.ValAddress, pubKey crypto.PubKey, description stakingtypes.Description, owner sdk.AccAddress) MsgCreateValidatorPOA {
	return MsgCreateValidatorPOA{
		Name:        name,
		Address:     address,
		PubKey:      pubKey,
		Description: description,
		Owner:       owner,
	}
}

// MarshalJSON implements the json.Marshaler interface to provide custom JSON
// serialization of the MsgCreateValidator type.
// We define a custom marshaler here to allow for msg to be used in the genesis file
func (msg MsgCreateValidatorPOA) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgCreateValidatorJSON{
		Name:        msg.Name,
		Address:     msg.Address,
		PubKey:      sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, msg.PubKey),
		Description: msg.Description,
		Owner:       msg.Owner,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface to provide custom
// JSON deserialization of the MsgCreateValidatorPOA type.
func (msg *MsgCreateValidatorPOA) UnmarshalJSON(bz []byte) error {
	var msgCreateValJSON msgCreateValidatorJSON
	if err := json.Unmarshal(bz, &msgCreateValJSON); err != nil {
		return err
	}

	msg.Name = msgCreateValJSON.Name
	msg.Address = msgCreateValJSON.Address
	var err error
	msg.PubKey, err = sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, msgCreateValJSON.PubKey)
	if err != nil {
		return err
	}
	msg.Description = msgCreateValJSON.Description
	msg.Owner = msgCreateValJSON.Owner

	return nil
}

// Route should return the name of the module
func (msg MsgCreateValidatorPOA) Route() string { return "poa" }

// Type should return the action
func (msg MsgCreateValidatorPOA) Type() string { return "create_validator_poa" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateValidatorPOA) ValidateBasic() error {
	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty name")
	}
	if msg.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address not specified")
	}
	if msg.Description == (stakingtypes.Description{}) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty description")
	}

	return nil
}

// GetSigners defines whose signature is required
func (msg MsgCreateValidatorPOA) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateValidatorPOA) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgCreateValidatorPOA{}, "poa/MsgCreateValidatorPOA", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
