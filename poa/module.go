package poa

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/allinbits/modules/poa/client/cli"
	"github.com/allinbits/modules/poa/client/rest"
	"github.com/allinbits/modules/poa/keeper"
	"github.com/allinbits/modules/poa/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

const (
	flagPubKey          = "pubkey"
	flagNodeId          = "node-id"
	flagIP              = "ip"
	flagMoniker         = "moniker"
	flagIdentity        = "identity"
	flagWebsite         = "website"
	flagSecurityContact = "security-contact"
	flagDetails         = "details"
)

// Type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the poa module.
type AppModuleBasic struct{}

// Name returns the poa module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterCodec registers the poa module's types for the given codec.
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the poa
// module.
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the poa module.
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data types.GenesisState
	err := ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	return types.ValidateGenesis(data)
}

// RegisterRESTRoutes registers the REST routes for the poa module.
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr)
}

// GetTxCmd returns the root tx command for the poa module.
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(cdc)
}

// GetQueryCmd returns no root query command for the poa module.
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey, cdc)
}

// extra helpers - gen-tx
// NOTE: abstact these functions to their own file

//// CreateValidatorMsgHelpers - used for gen-tx
func (AppModuleBasic) CreateValidatorMsgHelpers(ipDefault string) (
	fs *flag.FlagSet, nodeIDFlag, pubkeyFlag, amountFlag, defaultsDesc string) {
	viper.Set(flagIP, ipDefault)

	fs = flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(flagMoniker, "", "The validator's name")
	fs.String(flagWebsite, "", "The validator's (optional) website")
	fs.String(flagSecurityContact, "", "The validator's (optional) security contact email")
	fs.String(flagDetails, "", "The validator's (optional) details")
	fs.String(flagIdentity, "", "The (optional) identity signature (ex. UPort or Keybase)")
	return fs, flagNodeId, flagPubKey, "nil", "nil"
}

//// PrepareFlagsForTxCreateValidator - used for gen-tx
func (AppModuleBasic) PrepareFlagsForTxCreateValidator(config *cfg.Config, nodeID,
	chainID string, valPubKey crypto.PubKey) {
	viper.Set(flags.FlagChainID, chainID)
	viper.Set(flags.FlagFrom, viper.GetString(flags.FlagName))
	viper.Set(flagPubKey, sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, valPubKey))
	viper.Set(flagNodeId, nodeID)
}

//// BuildCreateValidatorMsg - used for gen-tx
func (AppModuleBasic) BuildCreateValidatorMsg(cliCtx context.CLIContext,
	txBldr authtypes.TxBuilder) (authtypes.TxBuilder, sdk.Msg, error) {
	pkStr := viper.GetString(flagPubKey)

	valAddr := cliCtx.GetFromAddress()
	consAddr := sdk.ValAddress(cliCtx.GetFromAddress())

	pk, _ := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, pkStr)

	moniker := viper.GetString(flagMoniker)
	identity := viper.GetString(flagIdentity)
	website := viper.GetString(flagWebsite)
	securityContact := viper.GetString(flagSecurityContact)
	details := viper.GetString(flagDetails)

	msg := types.NewMsgCreateValidatorPOA(
		consAddr.String(),
		consAddr,
		pk,
		stakingtypes.NewDescription(moniker, identity, website, securityContact, details),
		valAddr,
	)
	ip := viper.GetString(flagIP)
	nodeID := viper.GetString(flagNodeId)

	txBldr = txBldr.WithMemo(fmt.Sprintf("%s@%s:26656", nodeID, ip))

	return txBldr, msg, nil
}

//____________________________________________________________________________

// AppModule implements an application module for the poa module.
type AppModule struct {
	AppModuleBasic

	keeper     keeper.Keeper
	coinKeeper bank.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(k keeper.Keeper, bankKeeper bank.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
		coinKeeper:     bankKeeper,
	}
}

// Name returns the poa module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants registers the poa module invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns the message routing key for the poa module.
func (AppModule) Route() string {
	return types.RouterKey
}

// NewHandler returns an sdk.Handler for the poa module.
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

// QuerierRoute returns the poa module's querier route name.
func (AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

// NewQuerierHandler returns the poa module sdk.Querier.
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

// InitGenesis performs genesis initialization for the poa module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	return InitGenesis(ctx, am.keeper, genesisState)
}

// ExportGenesis returns the exported genesis state as raw bytes for the poa
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return ModuleCdc.MustMarshalJSON(gs)
}

// BeginBlock returns the begin blocker for the poa module.
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	BeginBlocker(ctx, req, am.keeper)
}

// EndBlock returns the end blocker for the poa module. It returns no validator
// updates.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return EndBlocker(ctx, am.keeper)
}
