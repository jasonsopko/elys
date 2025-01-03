package types

import (
	errorsmod "cosmossdk.io/errors"
)

// DONTCOVER

// x/parameter module sentinel errors
var (
	ErrInvalidMinCommissionRate    = errorsmod.Register(ModuleName, 1101, "invalid min commission rate")
	ErrInvalidMaxVotingPower       = errorsmod.Register(ModuleName, 1102, "invalid max voting power")
	ErrInvalidMinSelfDelegation    = errorsmod.Register(ModuleName, 1103, "invalid min self delegation")
	ErrInvalidRewardsDataLifecycle = errorsmod.Register(ModuleName, 1104, "invalid rewards data lifecycle")
)
