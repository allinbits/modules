package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/modules/poa/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group poa queries under a subcommand
	poaQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	poaQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdValidatorName(queryRoute, cdc),
			GetCmdValidatorsAll(queryRoute, cdc),
			GetCmdVote(queryRoute, cdc),
			GetCmdVotesAll(queryRoute, cdc),
		)...,
	)

	return poaQueryCmd
}

func GetCmdValidatorName(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "validator-poa [name]",
		Short: "validator-poa name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]
			route := fmt.Sprintf("custom/%s/validator-poa/%s", queryRoute, name)

			bz, err := cdc.MarshalJSON(types.NewQueryValidatorParams(name))
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(route, bz)
			if err != nil {
				fmt.Printf("could not resolve name - %s \n", name)
				return nil
			}

			var out types.QueryValidatorResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdValidatorsAll(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "validators",
		Short: "validators",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			resKVs, _, err := cliCtx.QuerySubspace(types.ValidatorsKey, queryRoute)
			if err != nil {
				return err
			}

			var validators []types.Validator
			for _, kv := range resKVs {
				validator := types.Validator{}
				cdc.UnmarshalBinaryLengthPrefixed(kv.Value, &validator)
				validators = append(validators, validator)

			}

			return cliCtx.PrintOutput(validators)
		},
	}
}

func GetCmdVote(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "vote-poa [name] [voter]",
		Short: "vote-poa name voter",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]
			voter := args[1]
			route := fmt.Sprintf("custom/%s/vote-poa/%s", queryRoute, voter)

			bz, err := cdc.MarshalJSON(types.NewQueryVoteParams(name, voter))
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(route, bz)
			if err != nil {
				fmt.Printf("could not resolve name - %s \n", voter)
				return nil
			}

			var out types.QueryVoteResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdVotesAll(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "votes",
		Short: "votes",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			resKVs, _, err := cliCtx.QuerySubspace(types.VotesKey, queryRoute)
			if err != nil {
				return err
			}

			var votes []types.Vote
			for _, kv := range resKVs {
				vote := types.Vote{}
				cdc.UnmarshalBinaryLengthPrefixed(kv.Value, &vote)
				votes = append(votes, vote)

			}

			return cliCtx.PrintOutput(votes)
		},
	}
}
