package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Resource defines the resource structure.
// Original Resource structure:
type Resource struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}

// Breaking change:
type Resource struct {
	ID          int    `json:"id" yaml:"id"` // Changed from string to int
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}


// GenesisState defines the resource module's genesis state.
type GenesisState struct {
	Resources []Resource `json:"resources" yaml:"resources"`
}
