package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/allinbits/modules/poa/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/poa/votes",
		getAllVotesHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/poa/validators",
		getAllValidatorsHandlerFn(cliCtx),
	).Methods("GET")

}

// HTTP request handler to query a validator votes
func getAllVotesHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cdc := codec.New()
		cliCtx = cliCtx.WithCodec(cdc)

		resKVs, height, err := cliCtx.QuerySubspace(types.VotesKey, "poa")
		if err != nil {
			panic(err)
		}

		var votes []types.Vote
		for _, kv := range resKVs {
			vote := types.Vote{}
			cdc.UnmarshalBinaryLengthPrefixed(kv.Value, &vote)
			votes = append(votes, vote)

		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, votes)
	}
}

func getAllValidatorsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cdc := codec.New()
		codec.RegisterCrypto(cdc)
		cliCtx = cliCtx.WithCodec(cdc)

		resKVs, height, err := cliCtx.QuerySubspace(types.ValidatorsKey, "poa")
		if err != nil {
			panic(err)
		}

		var validators []types.Validator
		for _, kv := range resKVs {
			validator := types.Validator{}
			cdc.UnmarshalBinaryLengthPrefixed(kv.Value, &validator)
			validators = append(validators, validator)

		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, validators)
	}
}
