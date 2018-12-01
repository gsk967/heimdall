package staking

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hmCommon "github.com/maticnetwork/heimdall/common"
	"github.com/maticnetwork/heimdall/helper"
	hmTypes "github.com/maticnetwork/heimdall/types"
)

func NewHandler(k hmCommon.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgValidatorJoin:
			return handleMsgValidatorJoin(ctx, msg, k)
		case MsgValidatorExit:
			return handleMsgValidatorExit(ctx, msg, k)
		case MsgSignerUpdate:
			return handleMsgSignerUpdate(ctx, msg, k)
		default:
			return sdk.ErrTxDecode("Invalid message in checkpoint module").Result()
		}
	}
}

func handleMsgValidatorJoin(ctx sdk.Context, msg MsgValidatorJoin, k hmCommon.Keeper) sdk.Result {
	//fetch validator from mainchain
	validator, err := helper.GetValidatorInfo(msg.ValidatorAddress)
	if err != nil {
		return hmCommon.ErrNoValidator(k.Codespace).Result()
	}

	pubKey := helper.BytesToPubkey(msg.ValidatorPubKey)

	if !bytes.Equal(validator.Address.Bytes(), msg.ValidatorAddress.Bytes()) {
		// TODO revert if mainchain doesn't match with incoming data
	}

	var savedValidator hmTypes.Validator
	err = k.GetValidatorInfo(ctx, msg.ValidatorAddress.Bytes(), &savedValidator)
	if err == nil {
		return hmCommon.ErrValidatorAlreadyJoined(k.Codespace).Result()
	}
	validator.PubKey = helper.BytesToPubkey(msg.ValidatorPubKey)

	// add validator to store
	err = k.AddValidator(ctx, validator)
	if err != nil {
		return hmCommon.ErrValidatorSave(k.Codespace).Result()
	}

	// validator set changed
	k.SetValidatorSetChangedFlag(ctx, true)

	return sdk.Result{}
}

func handleMsgSignerUpdate(ctx sdk.Context, msg MsgSignerUpdate, k hmCommon.Keeper) sdk.Result {
	var validator hmTypes.Validator

	// pull val from store
	err := k.GetValidatorInfo(ctx, msg.ValidatorAddress.Bytes(), &validator)
	if err != nil {
		hmCommon.StakingLogger.Error("Fetching of validator from store failed", "error", err, "validatorAddress", msg.ValidatorAddress)
		return hmCommon.ErrNoValidator(k.Codespace).Result()
	}

	pubKey := helper.BytesToPubkey(msg.NewValidatorPubKey)

	// check for already updated
	if !bytes.Equal(pubKey.Address().Bytes(), validator.Signer.Bytes()) {
		hmCommon.StakingLogger.Error("No signer update on stakemanager found or signer already updated", "error", err, "currentSigner", validator.Signer.String(), "signerFromMsg", pubKey.Address().String())
		return hmCommon.ErrValidatorAlreadySynced(k.Codespace).Result()
	}

	// update
	err = k.UpdateSigner(ctx, common.HexToAddress(pubKey.Address().String()), pubKey, msg.ValidatorAddress)
	if err != nil {
		hmCommon.StakingLogger.Error("Unable to update signer", "Error", err, "currentSigner", validator.Signer.String(), "signerFromMsg", pubKey.Address().String())
		panic(err)
	}

	// TODO: make prev signer power 0 , add new validator with new signer with same power and validator addr

	return sdk.Result{}
}

func handleMsgValidatorExit(ctx sdk.Context, msg MsgValidatorExit, k hmCommon.Keeper) sdk.Result {
	var validator hmTypes.Validator

	// fetch validator from store
	err := k.GetValidatorInfo(ctx, msg.ValidatorAddress.Bytes(), &validator)
	if err != nil {
		hmCommon.StakingLogger.Error("Fetching of validator from store failed", "error", err, "validatorAddress", msg.ValidatorAddress)
		return hmCommon.ErrNoValidator(k.Codespace).Result()
	}

	// check if validator deactivation period is set
	if validator.EndEpoch != 0 {
		hmCommon.StakingLogger.Error("Validator already unbonded")
		return hmCommon.ErrValUnbonded(k.Codespace).Result()
	}

	// TODO make sure this isnt needed
	// means exit has been processed but validator in unbonding period
	//if validator.Power != 0 {
	//	hmCommon.StakingLogger.Error("Validator already unbonded")
	//	return hmCommon.ErrValUnbonded(k.Codespace).Result()
	//}

	// Add deactivation time for validator
	k.AddDeactivationEpoch(ctx, msg.ValidatorAddress, validator)

	return sdk.Result{}
}
