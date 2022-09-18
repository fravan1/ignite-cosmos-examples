package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/nameservice module sentinel errors
var (
	ErrInsufficientFunds = sdkerrors.Register(ModuleName, 1000, "insufficient funds")
	ErrPriceParsing      = sdkerrors.Register(ModuleName, 1101, "price parsing error")
	ErrBidParsing        = sdkerrors.Register(ModuleName, 1102, "bid parsing error")
	ErrConvertingOwner   = sdkerrors.Register(ModuleName, 1201, "converting owner error")
	ErrConvertingBuyer   = sdkerrors.Register(ModuleName, 1202, "converting buyer error")
	ErrUnauthorized      = sdkerrors.Register(ModuleName, 1300, "unauthorized operation")
	ErrInvalidRequest    = sdkerrors.Register(ModuleName, 1401, "invalid request")
)
