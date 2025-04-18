package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/amm/types"
)

func (k Keeper) ExitPool(
	ctx sdk.Context,
	sender sdk.AccAddress,
	poolId uint64,
	shareInAmount math.Int,
	tokenOutMins sdk.Coins,
	tokenOutDenom string,
	isLiquidation, applyWeightBreakingFee bool,
) (exitCoins sdk.Coins, weightBalanceBonus math.LegacyDec, swapFee math.LegacyDec, slippage math.LegacyDec, takerFeesFinal math.LegacyDec, err error) {
	pool, poolExists := k.GetPool(ctx, poolId)
	if !poolExists {
		return sdk.Coins{}, math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), types.ErrInvalidPoolId
	}

	totalSharesAmount := pool.GetTotalShares()
	if shareInAmount.GTE(totalSharesAmount.Amount) {
		return sdk.Coins{}, math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), errorsmod.Wrapf(types.ErrInvalidMathApprox, "Trying to exit >= the number of shares contained in the pool.")
	} else if shareInAmount.LTE(math.ZeroInt()) {
		return sdk.Coins{}, math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), errorsmod.Wrapf(types.ErrInvalidMathApprox, "Trying to exit a negative amount of shares")
	}
	params := k.GetParams(ctx)
	takersFees := k.parameterKeeper.GetParams(ctx).TakerFees
	exitCoins, weightBalanceBonus, slippage, swapFee, takerFeesFinal, slippageCoins, err := pool.ExitPool(ctx, k.oracleKeeper, k.accountedPoolKeeper, shareInAmount, tokenOutDenom, params, takersFees, applyWeightBreakingFee)
	if err != nil {
		return sdk.Coins{}, math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), err
	}
	// Check treasury and update weightBalance
	if weightBalanceBonus.IsPositive() && exitCoins.Len() == 1 {
		rebalanceTreasuryAddr := sdk.MustAccAddressFromBech32(pool.GetRebalanceTreasury())
		treasuryTokenAmount := k.bankKeeper.GetBalance(ctx, rebalanceTreasuryAddr, exitCoins[0].Denom).Amount

		bonusTokenAmount := exitCoins[0].Amount.ToLegacyDec().Mul(weightBalanceBonus).TruncateInt()

		if treasuryTokenAmount.LT(bonusTokenAmount) {
			weightBalanceBonus = treasuryTokenAmount.ToLegacyDec().Quo(exitCoins[0].Amount.ToLegacyDec())
		}
	}

	if pool.PoolParams.UseOracle {
		if slippageCoins.IsAllPositive() {
			k.TrackWeightBreakingSlippage(ctx, pool.PoolId, sdk.NewCoin(slippageCoins[0].Denom, slippageCoins[0].Amount))
		}
	}

	if !tokenOutMins.DenomsSubsetOf(exitCoins) || tokenOutMins.IsAnyGT(exitCoins) {
		return sdk.Coins{}, math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), errorsmod.Wrapf(types.ErrLimitMinAmount,
			"Exit pool returned %s , minimum tokens out specified as %s",
			exitCoins, tokenOutMins)
	}

	err = k.ApplyExitPoolStateChange(ctx, pool, sender, shareInAmount, exitCoins, isLiquidation, weightBalanceBonus, takerFeesFinal, swapFee, slippageCoins)
	if err != nil {
		return sdk.Coins{}, math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), err
	}

	err = k.RecordTotalLiquidityDecrease(ctx, exitCoins)
	if err != nil {
		return sdk.Coins{}, math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), err
	}

	return exitCoins, weightBalanceBonus, swapFee, slippage, takerFeesFinal, nil
}
