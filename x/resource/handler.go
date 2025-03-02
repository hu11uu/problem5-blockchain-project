package resource

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"x/myblockchain/x/resource/types"
)

func handleMsgCreateResource(ctx sdk.Context, keeper Keeper, msg types.MsgCreateResource) (*sdk.Result, error) {
	resource := types.Resource{
		ID:          msg.Name, // In this example, using `Name` as `ID` (you can use a unique ID generation logic)
		Name:        msg.Name,
		Description: msg.Description,
	}

	// Save the resource
	keeper.SetResource(ctx, resource)

	return &sdk.Result{}, nil
}


func handleMsgUpdateResource(ctx sdk.Context, keeper Keeper, msg types.MsgUpdateResource) (*sdk.Result, error) {
	resource, found := keeper.GetResource(ctx, msg.ID)
	if !found {
		return nil, fmt.Errorf("resource not found")
	}

	resource.Name = msg.Name
	resource.Description = msg.Description

	// Update the resource
	keeper.SetResource(ctx, resource)

	return &sdk.Result{}, nil
}


func handleMsgDeleteResource(ctx sdk.Context, keeper Keeper, msg types.MsgDeleteResource) (*sdk.Result, error) {
	_, found := keeper.GetResource(ctx, msg.ID)
	if !found {
		return nil, fmt.Errorf("resource not found")
	}

	// Delete the resource
	keeper.DeleteResource(ctx, msg.ID)

	return &sdk.Result{}, nil
}
s