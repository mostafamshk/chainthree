package chainthree_test

import (
	"testing"

	keepertest "chainthree/testutil/keeper"
	"chainthree/testutil/nullify"
	"chainthree/x/chainthree"
	"chainthree/x/chainthree/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChainthreeKeeper(t)
	chainthree.InitGenesis(ctx, *k, genesisState)
	got := chainthree.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
