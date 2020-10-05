package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: change name to address of validator

// Vote - Contains data that allows validators to vote
type Vote struct {
	Voter   sdk.ValAddress `json:"voter" yaml:"voter"`
	Name    string         `json:"name" yaml:"name"`
	InFavor bool           `json:"in_favor" yaml:"in_favor"`
}

// NewVote
func NewVote(voter sdk.ValAddress, name string, inFavor bool) Vote {
	return Vote{
		Voter:   voter,
		Name:    name,
		InFavor: inFavor,
	}
}
