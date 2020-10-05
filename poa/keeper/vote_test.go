package keeper

import (
	"testing"

	"github.com/allinbits/modules/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeeperVoteFunctions(t *testing.T) {
	ctx, keeper := MakeTestCtxAndKeeper(t)

	// Create test vote
	valPubKey := MakeTestPubKey(SamplePubKey)
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())

	valName := "name"

	vote := types.NewVote(
		valAddr,
		valName,
		true,
	)

	// TODO: split into multiple test cases

	// Set a vote in the store
	keeper.SetVote(ctx, vote)

	// Check the store to see if the vote was saved
	retVal, found := keeper.GetVote(ctx, append([]byte(valName), valAddr...))
	assert.Equal(t, vote, retVal, "Should return the correct vote from the store")
	require.True(t, found)

	// Get all votes from the store
	allVals := keeper.GetAllVotes(ctx)
	require.Equal(t, 1, len(allVals))

	// Get all votes for a validator from the store
	votesForVal := keeper.GetAllVotesForValidator(ctx, valName)
	require.Equal(t, 1, len(votesForVal))
}
