package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/allinbits/modules/poa/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	poaTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	poaTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreateValidator(cdc),
		GetCmdVoteValidator(cdc, "vote-validator", true),
		GetCmdVoteValidator(cdc, "kick-validator", false),
	)...)

	return poaTxCmd
}

// GetCmdCreateValidator is the CLI command for sending a CreateValidator transaction
func GetCmdCreateValidator(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-validator [name] [public key] [moniker] [identity] [website] [security_contract] [details]",
		Short: "create a validatator with a name and public key",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			accAddr := cliCtx.GetFromAddress()
			consAddr := sdk.ValAddress(accAddr)
			valPubKey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateValidatorPOA(
				args[0],
				consAddr,
				valPubKey,
				stakingtypes.NewDescription(args[2], args[3], args[4], args[5], args[6]),
				accAddr,
			)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdVoteValidator is the CLI command for sending a VoteValidator transaction
func GetCmdVoteValidator(cdc *codec.Codec, use string, inFavor bool) *cobra.Command {
	return &cobra.Command{
		Use:   use + " [validator-name]",
		Short: "vote for a validator using their name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			accAddr := cliCtx.GetFromAddress()
			valAddr := sdk.ValAddress(accAddr)

			msg := types.NewMsgVoteValidator(args[0], valAddr, inFavor, accAddr)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
