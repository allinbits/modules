package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestVoteEqual(t *testing.T) {
	valPubKey := MakeTestPubKey(SamplePubKey)
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())

	vote1 := NewVote(
		valAddr,
		"name",
		true,
	)

	vote2 := NewVote(
		valAddr,
		"name",
		true,
	)

	// Test two validator are equal when created with the same data
	require.True(t, vote1.Voter.Equals(vote2.Voter))

	valPubKey3 := MakeTestPubKey(SamplePubKey2)
	valAddr3 := sdk.ValAddress(valPubKey3.Address().Bytes())
	vote3 := NewVote(
		valAddr3,
		"name",
		false,
	)

	// Test two votes are no equal when created with different data
	require.False(t, vote1.Voter.Equals(vote3.Voter))
}
