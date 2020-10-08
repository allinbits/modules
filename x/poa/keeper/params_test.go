package keeper

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/allinbits/modules/poa/types"
)

func TestKeeperParamsFunctions(t *testing.T) {
	ctx, keeper := MakeTestCtxAndKeeper(t)

	// SetParams test
	keeper.SetParams(ctx, types.DefaultParams())

	// GetParams test
	params := keeper.GetParams(ctx)

	require.Equal(t, uint16(49), params.Quorum)
	require.Equal(t, uint16(100), params.MaxValidators)
}
