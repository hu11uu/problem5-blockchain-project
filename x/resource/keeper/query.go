package keeper

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	"x/myblockchain/x/resource/types"
)

func (k Keeper) QueryResources(ctx types.Context) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	var resources []types.Resource

	iter := store.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var resource types.Resource
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &resource)
		resources = append(resources, resource)
	}

	res, err := json.Marshal(resources)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal resources: %v", err)
	}
	return res, nil
}
