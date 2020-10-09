package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/allinbits/modules/x/poa/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	"github.com/cosmos/cosmos-sdk/x/bank"
)

// SelectorFn allows an entity to be selected by certain conditions
type SelectorFn func(interface{}) bool

// UnmarshalFn is a generic function to unmarshal bytes
type UnmarshalFn func(value []byte) (interface{}, bool)

// Keeper of the poa store
type Keeper struct {
	CoinKeeper bank.Keeper
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	paramspace types.ParamSubspace
}

// NewKeeper creates a poa keeper
func NewKeeper(coinKeeper bank.Keeper, cdc *codec.Codec, key sdk.StoreKey, paramspace params.Subspace) Keeper {
	keeper := Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   key,
		cdc:        cdc,
		paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Set sets a value in the db with a prefixed key
func (k Keeper) Set(ctx sdk.Context, key []byte, prefix []byte, i interface{}) {
	store := ctx.KVStore(k.storeKey)
	store.Set(append(prefix, key...), k.cdc.MustMarshalBinaryLengthPrefixed(i))
}

func (k Keeper) Delete(ctx sdk.Context, key []byte, prefix []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(prefix, key...))
}

// Get gets an item from the store by bytes
func (k Keeper) Get(ctx sdk.Context, key []byte, prefix []byte, unmarshal UnmarshalFn) (i interface{}, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(append(prefix, key...))

	return unmarshal(value)
}

// GetAll values from with a prefix from the store
func (k Keeper) GetAll(ctx sdk.Context, prefix []byte, unmarshal UnmarshalFn) (i []interface{}) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		value, _ := unmarshal(iterator.Value())
		i = append(i, value)
	}
	return i
}
