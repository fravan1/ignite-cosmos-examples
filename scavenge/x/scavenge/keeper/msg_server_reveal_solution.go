package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"crypto/sha256"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	"scavenge/x/scavenge/types"
)

func (k msgServer) RevealSolution(goCtx context.Context, msg *types.MsgRevealSolution) (*types.MsgRevealSolutionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var solutionScavengerBytes = []byte(msg.Solution + msg.Creator)

	var solutionScavengerHash = sha256.Sum256(solutionScavengerBytes)

	var solutionScavengerHashString = hex.EncodeToString(solutionScavengerHash[:])

	_, isFound := k.GetCommit(ctx, solutionScavengerHashString)

	if !isFound {
		return nil, errors.Wrap(types.ErrInvalidRequest, "Commit with that hash doesn't exists")
	}

	var solutionHash = sha256.Sum256([]byte(msg.Solution))

	var solutionHashString = hex.EncodeToString(solutionHash[:])
	var scavenge types.Scavenge

	scavenge, isFound = k.GetScavenge(ctx, solutionHashString)

	if !isFound {
		return nil, errors.Wrap(types.ErrInvalidRequest, "Scavenge with that solution hash doesn't exists")
	}

	_, err := sdk.AccAddressFromBech32(scavenge.Scavenger)

	if err == nil {
		return nil, errors.Wrap(types.ErrInvalidRequest, "Scavenge has already been solved")
	}

	scavenge.Scavenger = msg.Creator

	scavenge.Solution = msg.Solution

	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	scavenger, err := sdk.AccAddressFromBech32(scavenge.Scavenger)
	if err != nil {
		panic(err)
	}

	reward, err := sdk.ParseCoinsNormalized(scavenge.Reward)
	if err != nil {
		panic(err)
	}

	sdkError := k.bankKeeper.SendCoins(ctx, moduleAcct, scavenger, reward)
	if sdkError != nil {
		return nil, sdkError
	}

	k.SetScavenge(ctx, scavenge)
	return &types.MsgRevealSolutionResponse{}, nil
}
