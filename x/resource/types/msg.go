package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// MsgCreateResource defines a message to create a resource.
type MsgCreateResource struct {
	Creator     string `json:"creator" yaml:"creator"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}

func NewMsgCreateResource(creator, name, description string) MsgCreateResource {
	return MsgCreateResource{
		Creator:     creator,
		Name:        name,
		Description: description,
	}
}

// Route should return the name of the module.
func (msg MsgCreateResource) Route() string { return RouterKey }

// Type should return the action.
func (msg MsgCreateResource) Type() string { return "create_resource" }

// GetSigners defines whose signature is required.
func (msg MsgCreateResource) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

// MsgUpdateResource defines a message to update an existing resource.
type MsgUpdateResource struct {
	Creator     string `json:"creator" yaml:"creator"`
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}

// MsgDeleteResource defines a message to delete a resource.
type MsgDeleteResource struct {
	Creator string `json:"creator" yaml:"creator"`
	ID      string `json:"id" yaml:"id"`
}
