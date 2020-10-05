package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all authority state that must be provided at genesis
type GenesisState struct {
	Params     Params      `json:"params" yaml:"params"`
	Validators []Validator `json:"validators" yaml:"validators"`
}

// LastValidatorPower required for validator set update logic
type LastValidatorPower struct {
	Address sdk.ValAddress
	Power   int64
}

// NewGenesisState creates a new GenesisState instanc e
func NewGenesisState(params Params, validators []Validator) GenesisState {
	return GenesisState{
		Params:     params,
		Validators: validators,
	}
}

// DefaultGenesisState gets the raw genesis raw message for testing
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params: DefaultParams(),
	}
}

// ValidateGenesis validates the authority genesis parameters
func ValidateGenesis(data GenesisState) error {
	// TODO: Create a sanity check to make sure the state conforms to the modules needs
	return nil
}
