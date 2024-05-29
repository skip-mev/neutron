package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/dex module sentinel errors

var (
	ErrInvalidTradingPair = sdkerrors.Register(
		ModuleName,
		1102,
		"Invalid token pair:",
	) // "%s<>%s", tokenA, tokenB
	ErrInsufficientShares = sdkerrors.Register(
		ModuleName,
		1104,
		"Insufficient shares:",
	) // "%s does not have %s shares of type %s", address, shares, sharesID
	ErrUnbalancedTxArray = sdkerrors.Register(
		ModuleName,
		1110,
		"Transaction input arrays are not of the same length.",
	)
	ErrValidLimitOrderTrancheNotFound = sdkerrors.Register(
		ModuleName,
		1111,
		"Limit order trache not found:",
	) // "%d", trancheKey
	ErrCancelEmptyLimitOrder = sdkerrors.Register(
		ModuleName,
		1112,
		"Cannot cancel additional liquidity from limit order tranche:",
	) // "%d", tranche.TrancheKey
	ErrInsufficientLiquidity = sdkerrors.Register(
		ModuleName,
		1114,
		"Not enough liquidity to complete trade",
	)
	ErrTickOutsideRange = sdkerrors.Register(
		ModuleName,
		1117,
		"abs(tick) + fee must be < 559,680",
	)
	ErrInvalidPoolDenom = sdkerrors.Register(
		ModuleName,
		1118,
		"Denom is not an instance of Neutron PoolDenom",
	)
	ErrInvalidPairIDStr = sdkerrors.Register(
		ModuleName,
		1119,
		"PairID does not conform to pattern TokenA<>TokenB",
	)
	ErrZeroDeposit = sdkerrors.Register(
		ModuleName,
		1120,
		"At least one deposit amount must be > 0.",
	)
	ErrZeroTrueDeposit = sdkerrors.Register(
		ModuleName,
		1121,
		"Cannot deposit double-sided liquidity in tick with prexisting single-sided liquidity.",
	)
	ErrWithdrawEmptyLimitOrder = sdkerrors.Register(
		ModuleName,
		1124,
		"Cannot withdraw additional liqudity from this limit order at this time.",
	)
	ErrZeroSwap = sdkerrors.Register(
		ModuleName,
		1125,
		"MaxAmountIn in must be > 0 for swap.",
	)
	ErrActiveLimitOrderNotFound = sdkerrors.Register(
		ModuleName,
		1128,
		"No active limit found. It does not exist or has already been filled",
	)
	ErrZeroWithdraw = sdkerrors.Register(
		ModuleName,
		1129,
		"Withdraw amount must be > 0.",
	)
	ErrZeroLimitOrder = sdkerrors.Register(
		ModuleName,
		1130,
		"Limit order amount must be > 0.",
	)
	ErrDepositShareUnderflow = sdkerrors.Register(
		ModuleName,
		1133,
		"Deposit amount is too small to issue shares",
	)
	ErrFoKLimitOrderNotFilled = sdkerrors.Register(
		ModuleName,
		1134,
		"Fill Or Kill limit order couldn't be executed in its entirety.",
	)
	ErrInvalidTimeString = sdkerrors.Register(
		ModuleName,
		1135,
		"Time string must be formatted as MM/dd/yyyy HH:mm:ss (ex. 02/05/2023 15:34:56) ",
	)
	ErrGoodTilOrderWithoutExpiration = sdkerrors.Register(
		ModuleName,
		1136,
		"Limit orders of type GOOD_TIL_TIME must supply an ExpirationTime.",
	)
	ErrExpirationOnWrongOrderType = sdkerrors.Register(
		ModuleName,
		1137,
		"Only Limit orders of type GOOD_TIL_TIME can supply an ExpirationTime.",
	)
	ErrInvalidOrderType = sdkerrors.Register(
		ModuleName,
		1138,
		"Order type must be one of: GOOD_TIL_CANCELLED, FILL_OR_KILL, IMMEDIATE_OR_CANCEL, JUST_IN_TIME, or GOOD_TIL_TIME.",
	)
	ErrExpirationTimeInPast = sdkerrors.Register(
		ModuleName,
		1139,
		"Limit order expiration time must be greater than current block time:",
	)
	ErrExitLimitPriceHit = sdkerrors.Register(
		ModuleName,
		1140,
		"ExitLimitPrice cannot be satisfied.",
	)
	ErrAllMultiHopRoutesFailed = sdkerrors.Register(
		ModuleName,
		1141,
		"All multihop routes failed limitPrice check or had insufficient liquidity",
	)
	ErrMultihopExitTokensMismatch = sdkerrors.Register(
		ModuleName,
		1142,
		"All multihop routes must have the same exit token",
	)
	ErrMissingMultihopRoute = sdkerrors.Register(
		ModuleName,
		1143,
		"Must supply at least 1 route for multihop swap",
	)
	ErrZeroMaxAmountOut = sdkerrors.Register(
		ModuleName,
		1144,
		"MaxAmountOut must be nil or > 0.",
	)
	ErrInvalidMaxAmountOutForMaker = sdkerrors.Register(
		ModuleName,
		1145,
		"MaxAmountOut can only be set for taker only limit orders.",
	)
	ErrInvalidFee = sdkerrors.Register(
		ModuleName,
		1148,
		"Fee must must a legal fee amount:",
	)
	ErrInvalidAddress = sdkerrors.Register(
		ModuleName,
		1149,
		"Invalid Address",
	)
	ErrRouteWithoutExitToken = sdkerrors.Register(
		ModuleName,
		1150,
		"Each route should specify at least two hops - input and output tokens",
	)
	ErrCycleInHops = sdkerrors.Register(
		ModuleName,
		1151,
		"Hops cannot have cycles",
	)
	ErrZeroExitPrice = sdkerrors.Register(
		ModuleName,
		1152,
		"Cannot have negative or zero exit price",
	)
	ErrDuplicatePoolDeposit = sdkerrors.Register(
		ModuleName,
		1153,
		"Can only provide a single deposit amount for each tick, fee pair",
	)
	ErrSwapAmountTooSmall = sdkerrors.Register(
		ModuleName,
		1154,
		"Swap amount too small; creates unfair spread for liquidity providers",
	)
	ErrDexPaused = sdkerrors.Register(
		ModuleName,
		1155,
		"Dex has been paused, all messages are disabled at this time",
	)
	ErrOverJITPerBlockLimit = sdkerrors.Register(
		ModuleName,
		1156,
		"Maximum JIT LimitOrders per block has already been reached",
	)
)
