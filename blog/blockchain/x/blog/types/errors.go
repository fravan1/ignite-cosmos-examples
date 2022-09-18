package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/blog module sentinel errors
var (
	ErrSample     = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrCommentOld = sdkerrors.Register(ModuleName, 1300, "")
	ErrID         = sdkerrors.Register(ModuleName, 1400, "")
)
