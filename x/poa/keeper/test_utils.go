package keeper

import (
	"bytes"
	"encoding/hex"

	"github.com/allinbits/modules/poa/types"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	dbm "github.com/tendermint/tm-db"
)

const (
	SamplePubKey  = "b7a3c12dc0c8c748ab07525b701122b88bd78f600c76342d27f25e5f92444cde"
	SamplePubKey2 = "b7a3c12dc0c8c748ab07525b701122b88bd78f600c76342d27f25e5f92444cdf"
)

func MakeTestPubKey(pk string) crypto.PubKey {
	var buffer bytes.Buffer
	buffer.WriteString(pk)

	pkBytes, err := hex.DecodeString(buffer.String())
	if err != nil {
		panic(err)
	}
	var pkEd ed25519.PubKeyEd25519
	copy(pkEd[:], pkBytes)
	return pkEd
}

func MakeTestCtxAndKeeper(t *testing.T) (sdk.Context, Keeper) {
	var cdc = codec.New()
	codec.RegisterCrypto(cdc)

	keyPoa := sdk.NewKVStoreKey(types.StoreKey)
	keyAcc := sdk.NewKVStoreKey(auth.StoreKey)
	keyParams := sdk.NewKVStoreKey(params.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(params.TStoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyPoa, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, abci.Header{ChainID: "foochainid"}, true, nil)

	pk := params.NewKeeper(
		cdc,
		keyParams,
		tkeyParams,
	)

	accountKeeper := auth.NewAccountKeeper(
		cdc,
		keyAcc,
		pk.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)

	bk := bank.NewBaseKeeper(
		accountKeeper,
		pk.Subspace(bank.DefaultParamspace),
		nil,
	)

	keeper := NewKeeper(
		bk,
		cdc,
		keyPoa,
		pk.Subspace(DefaultParamspace),
	)
	keeper.SetParams(ctx, types.DefaultParams())

	return ctx, keeper
}
