package keeper

import (
	"github.com/cosmos/cosmos-sdk/types"
	"x/myblockchain/x/resource/types"
)

type Keeper struct {
	storeKey types.StoreKey
	cdc      types.Codec
}

func (k Keeper) SetResource(ctx types.Context, resource types.Resource) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryBare(&resource)
	store.Set([]byte(resource.ID), b)
}

func (k Keeper) GetResource(ctx types.Context, id string) (types.Resource, bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get([]byte(id))
	if b == nil {
		return types.Resource{}, false
	}
	var resource types.Resource
	k.cdc.MustUnmarshalBinaryBare(b, &resource)
	return resource, true
}

func (k Keeper) DeleteResource(ctx types.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(id))
}
