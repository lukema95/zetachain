package types

import (
	sdkmath "cosmossdk.io/math"
)

// GetRewardsDistributions returns the current distribution of rewards
// for validators, observers and TSS signers
// If the percentage is not set, it returns 0
func GetRewardsDistributions(params Params) (sdkmath.Int, sdkmath.Int, sdkmath.Int) {
	// We do not need to validate params here,
	// as we can assume that the params have been validated before setting the values.
	blockReward := params.BlockRewardAmount

	// Fetch the validator rewards, use 0 if the percentage is not set
	validatorRewards := sdkmath.NewInt(0)
	validatorRewardsDec, err := sdkmath.LegacyNewDecFromStr(params.ValidatorEmissionPercentage)
	if err == nil {
		validatorRewards = validatorRewardsDec.Mul(blockReward).TruncateInt()
	}

	// Fetch the observer rewards, use 0 if the percentage is not set
	observerRewards := sdkmath.NewInt(0)
	observerRewardsDec, err := sdkmath.LegacyNewDecFromStr(params.ObserverEmissionPercentage)
	if err == nil {
		observerRewards = observerRewardsDec.Mul(blockReward).TruncateInt()
	}

	// Fetch the TSS signer rewards, use 0 if the percentage is not set
	tssSignerRewards := sdkmath.NewInt(0)
	tssSignerRewardsDec, err := sdkmath.LegacyNewDecFromStr(params.TssSignerEmissionPercentage)
	if err == nil {
		tssSignerRewards = tssSignerRewardsDec.Mul(blockReward).TruncateInt()
	}

	return validatorRewards, observerRewards, tssSignerRewards
}
