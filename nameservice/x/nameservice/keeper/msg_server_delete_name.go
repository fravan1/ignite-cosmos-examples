package keeper

import (
	"context"
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"nameservice/x/nameservice/types"
)

func (k msgServer) DeleteName(goCtx context.Context, msg *types.MsgDeleteName) (*types.MsgDeleteNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	whois, isFound := k.GetWhois(ctx, msg.Name)

	if !isFound {
		return nil, errors.Wrap(types.ErrInvalidRequest, "name doesn't exist")
	}

	if msg.Creator != whois.Owner {
		return nil, errors.Wrap(types.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWhois(ctx, msg.Name)

	return &types.MsgDeleteNameResponse{}, nil
}
