package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/params"
)

// Default parameter namespace
const (
	// Default percentage of votes to join the set
	DefaultQuorum uint16 = 49

	// Default maximum number of bonded validators
	DefaultMaxValidators uint16 = 100
)

// nolint - Keys for parameter access
var (
	KeyQuorum        = []byte("Quorum")
	KeyMaxValidators = []byte("MaxValidators")
)

// ParamKeyTable for poa module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// Params - used for initializing default parameter for poa at genesis
type Params struct {
	Quorum        uint16 `json:"quorum" yaml:"quorum"`                 // percentage of validators that need to vote
	MaxValidators uint16 `json:"max_validators" yaml:"max_validators"` // maximum number of validators (max uint16 = 65535)
}

// NewParams creates a new Params object
func NewParams(quorum uint16, maxValidators uint16) Params {
	return Params{
		Quorum:        quorum,
		MaxValidators: maxValidators,
	}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	return fmt.Sprintf(`
	Quorum: %v, MaxValidators: %v
	`, p.Quorum, p.MaxValidators)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(KeyQuorum, &p.Quorum, validateQuorum),
		params.NewParamSetPair(KeyMaxValidators, &p.MaxValidators, validateMaxValidators),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(DefaultQuorum, DefaultMaxValidators)
}

func validateQuorum(i interface{}) error {
	val, ok := i.(uint16)
	if !ok {
		return fmt.Errorf("invalid type: %T", val)
	}

	if val > 100 {
		return fmt.Errorf("quorum must be less than 100: %d", val)
	}

	return nil
}

func validateMaxValidators(i interface{}) error {
	val, ok := i.(uint16)
	if !ok {
		return fmt.Errorf("invalid type: %T", i)
	}

	if val < 0 || val > 100 {
		return fmt.Errorf("max validators must between 0 and 100: %d", val)
	}

	return nil
}
