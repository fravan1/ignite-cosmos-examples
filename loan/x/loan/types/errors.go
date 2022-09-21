package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/loan module sentinel errors
var (
	ErrWrongLoanState = errors.Register(ModuleName, 2, "wrong loan state")
	ErrDeadline       = errors.Register(ModuleName, 3, "deadline")
)
