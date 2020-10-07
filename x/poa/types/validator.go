package types

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NOTE: Should we have a validator state? IsAccepted & InSet

// TODO: Change validator.Accepted => validator.IsAccepted

// Validator - Data structure to define validators for consensus
type Validator struct {
	Name        string                   `json:"name" yaml:"name"`
	Address     sdk.ValAddress           `json:"address" yaml:"address"`
	PubKey      crypto.PubKey            `json:"pubkey" yaml:"pubkey"`
	Description stakingtypes.Description `json:"description" yaml:"description"`
	Accepted    bool                     `json:"accepted" yaml:"accepted"`
	InSet       bool                     `json:"inset" yaml:"inset"`
}

// ABCIValidatorUpdate returns an abci.ValidatorUpdate from a staking validator type
// with the full validator power
func (v Validator) ABCIValidatorUpdate(power int64) abci.ValidatorUpdate {
	return abci.ValidatorUpdate{
		PubKey: tmtypes.TM2PB.PubKey(v.PubKey),
		Power:  power,
	}
}

// NewValidator - initialize a new validator
func NewValidator(name string, address sdk.ValAddress, pubKey crypto.PubKey, description stakingtypes.Description) Validator {
	return Validator{
		Name:        name,
		Accepted:    false,
		Address:     address,
		PubKey:      pubKey,
		Description: description,
	}
}
