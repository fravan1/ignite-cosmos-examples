package keeper

import (
	"context"
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"nameservice/x/nameservice/types"
)

func (k msgServer) SetName(goCtx context.Context, msg *types.MsgSetName) (*types.MsgSetNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	whois, _ := k.GetWhois(ctx, msg.Name)

	if msg.Creator != whois.Owner {
		return nil, errors.Wrap(types.ErrUnauthorized, "Incorrect owner")
	}

	newWhois := types.Whois{
		Index: msg.Name,
		Name:  msg.Name,
		Value: msg.Value,
		Owner: whois.Owner,
		Price: whois.Price,
	}

	k.SetWhois(ctx, newWhois)

	return &types.MsgSetNameResponse{}, nil
}
