package keeper

import (
	"context"
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"nameservice/x/nameservice/types"
)

func (k msgServer) BuyName(goCtx context.Context, msg *types.MsgBuyName) (*types.MsgBuyNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Try getting a name from the store
	whois, isFound := k.GetWhois(ctx, msg.Name)

	minPrice := sdk.Coins{sdk.NewInt64Coin("token", 10)}

	price, err := sdk.ParseCoinsNormalized(whois.Price)
	if err != nil {
		return nil, errors.Wrap(types.ErrPriceParsing, "price parsing failed")
	}
	bid, err := sdk.ParseCoinsNormalized(msg.Bid)
	if err != nil {
		return nil, errors.Wrap(types.ErrBidParsing, "bid parsing failed")
	}

	owner, _ := sdk.AccAddressFromBech32(whois.Owner)

	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)

	if isFound {
		if price.IsAllGT(bid) {
			return nil, errors.Wrap(types.ErrInsufficientFunds, "bid is not high enough")
		}

		// send tokens from the buyer to the owner
		k.bankKeeper.SendCoins(ctx, buyer, owner, bid)
	} else {
		if minPrice.IsAllGT(bid) {
			return nil, errors.Wrap(types.ErrInsufficientFunds, "bid is not high enough")
		}

		// send tokens from the buyer's account to the module's account
		k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, bid)
	}

	newWhois := types.Whois{
		Index: msg.Name,
		Name:  msg.Name,
		Value: whois.Value,
		Price: bid.String(),
		Owner: buyer.String(),
	}

	k.SetWhois(ctx, newWhois)

	return &types.MsgBuyNameResponse{}, nil
}
