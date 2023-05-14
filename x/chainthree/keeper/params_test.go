package keeper_test

import (
	"testing"

	testkeeper "chainthree/testutil/keeper"
	"chainthree/x/chainthree/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChainthreeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
