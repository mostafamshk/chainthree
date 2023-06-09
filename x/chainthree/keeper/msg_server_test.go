package keeper_test

import (
	"context"
	"testing"

	keepertest "chainthree/testutil/keeper"
	"chainthree/x/chainthree/keeper"
	"chainthree/x/chainthree/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChainthreeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
