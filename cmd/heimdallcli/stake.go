package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	cliContext "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/maticnetwork/bor/common"
	chainmanagerTypes "github.com/maticnetwork/heimdall/chainmanager/types"
	"github.com/maticnetwork/heimdall/helper"

	stakingcli "github.com/maticnetwork/heimdall/staking/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var checkpointEndpoint = "chainmanager/params"

// StakeCmd stakes for a validator
func StakeCmd(cliCtx cliContext.CLIContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake",
		Short: "Stake matic tokens for your account",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			helper.InitHeimdallConfig("")

			validatorStr := viper.GetString(stakingcli.FlagValidatorAddress)
			stakeAmountStr := viper.GetString(stakingcli.FlagAmount)
			feeAmountStr := viper.GetString(stakingcli.FlagFeeAmount)
			acceptDelegation := viper.GetBool(stakingcli.FlagAcceptDelegation)

			// validator str
			if validatorStr == "" {
				return errors.New("Validator address is required")
			}

			// stake amount
			stakeAmount, ok := big.NewInt(0).SetString(stakeAmountStr, 10)
			if !ok {
				return errors.New("Invalid stake amount")
			}

			// fee amount
			feeAmount, ok := big.NewInt(0).SetString(feeAmountStr, 10)
			if !ok {
				return errors.New("Invalid fee amount")
			}

			// contract caller
			contractCaller, err := helper.NewContractCaller()
			if err != nil {
				return err
			}

			route := fmt.Sprintf("custom/%s", checkpointEndpoint)

			bz, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var params chainmanagerTypes.Params
			json.Unmarshal(bz, &params)

			stakingManagerAddress := params.ChainParams.StakingManagerAddress.EthAddress()
			stakeManagerInstance, err := contractCaller.GetStakeManagerInstance(stakingManagerAddress)

			return contractCaller.StakeFor(
				common.HexToAddress(validatorStr),
				stakeAmount,
				feeAmount,
				acceptDelegation,
				stakingManagerAddress,
				stakeManagerInstance,
			)
		},
	}

	cmd.Flags().String(stakingcli.FlagValidatorAddress, "", "--validator=<validator address here>")
	cmd.Flags().String(stakingcli.FlagAmount, "10000000000000000000", "--staked-amount=<stake amount>, if left blank it will be assigned as 10 matic tokens")
	cmd.Flags().String(stakingcli.FlagFeeAmount, "5000000000000000000", "--fee-amount=<heimdall fee amount>, if left blank will be assigned as 5 matic tokens")
	cmd.Flags().Bool(stakingcli.FlagAcceptDelegation, true, "--accept-delegation=<accept delegation>, if left blank will be assigned as true")
	return cmd
}

// ApproveCmd approves tokens for a validator
func ApproveCmd(cliCtx cliContext.CLIContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve",
		Short: "Approve the tokens to stake",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			helper.InitHeimdallConfig("")

			stakeAmountStr := viper.GetString(stakingcli.FlagAmount)
			feeAmountStr := viper.GetString(stakingcli.FlagFeeAmount)

			// stake amount
			stakeAmount, ok := big.NewInt(0).SetString(stakeAmountStr, 10)
			if !ok {
				return errors.New("Invalid stake amount")
			}

			// fee amount
			feeAmount, ok := big.NewInt(0).SetString(feeAmountStr, 10)
			if !ok {
				return errors.New("Invalid fee amount")
			}

			contractCaller, err := helper.NewContractCaller()
			if err != nil {
				return err
			}
			route := fmt.Sprintf("custom/%s", checkpointEndpoint)

			bz, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var params chainmanagerTypes.Params
			json.Unmarshal(bz, &params)

			stakingManagerAddress := params.ChainParams.StakingManagerAddress.EthAddress()
			maticTokenAddress := params.ChainParams.MaticTokenAddress.EthAddress()

			maticTokenInstance, err := contractCaller.GetMaticTokenInstance(maticTokenAddress)

			return contractCaller.ApproveTokens(stakeAmount.Add(stakeAmount, feeAmount), stakingManagerAddress, maticTokenAddress, maticTokenInstance)
		},
	}

	cmd.Flags().String(stakingcli.FlagAmount, "10000000000000000000", "--staked-amount=<stake amount>, if left blank will be assigned as 10 matic tokens")
	cmd.Flags().String(stakingcli.FlagFeeAmount, "5000000000000000000", "--fee-amount=<heimdall fee amount>, if left blank will be assigned as 5 matic tokens")
	return cmd
}
