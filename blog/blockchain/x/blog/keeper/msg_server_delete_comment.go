package keeper

import (
	"context"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	"blog/x/blog/types"
	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteComment(goCtx context.Context, msg *types.MsgDeleteComment) (*types.MsgDeleteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	comment, exist := k.GetComment(ctx, msg.CommentID)
	if !exist {
		return nil, coserrors.Wrapf(types.ErrID, "Comment doesn't exist")
	}

	if msg.PostID != comment.PostID {
		return nil, coserrors.Wrapf(types.ErrID, "Post Blog ID does not exist for which comment with Blog Id %d was made", msg.PostID)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommentKey))
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, comment.Id)
	store.Delete(bz)

	return &types.MsgDeleteCommentResponse{}, nil
}
