package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/am3o/cosmos/x/cosmos/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for cosmos clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryListVote:
			return listVote(ctx, k)
		case types.QueryGetVote:
			return getVote(ctx, path[1:], k)
		case types.QueryListPoll:
			return listPoll(ctx, k)
		case types.QueryGetPoll:
			return getPoll(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown cosmos query endpoint")
		}
	}
}
