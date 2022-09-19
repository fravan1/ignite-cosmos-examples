package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/scavenge module sentinel errors
var (
	ErrInvalidRequest = sdkerrors.Register(ModuleName, 1100, "invalid request")
)
