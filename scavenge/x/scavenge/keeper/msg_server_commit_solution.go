package keeper

import (
	"context"
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"scavenge/x/scavenge/types"
)

func (k msgServer) CommitSolution(goCtx context.Context, msg *types.MsgCommitSolution) (*types.MsgCommitSolutionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	commit := types.Commit{
		Index:                 msg.SolutionScavengerHash,
		SolutionHash:          msg.SolutionHash,
		SolutionScavengerHash: msg.SolutionScavengerHash,
	}

	_, isFound := k.GetCommit(ctx, commit.SolutionScavengerHash)

	if isFound {
		return nil, errors.Wrap(types.ErrInvalidRequest, "Commit with that hash already exists")
	}

	k.SetCommit(ctx, commit)

	return &types.MsgCommitSolutionResponse{}, nil
}
