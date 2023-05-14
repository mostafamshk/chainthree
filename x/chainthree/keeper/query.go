package keeper

import (
	"chainthree/x/chainthree/types"
)

var _ types.QueryServer = Keeper{}
