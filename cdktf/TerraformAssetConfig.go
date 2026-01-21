// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cdktf


// Experimental.
type TerraformAssetConfig struct {
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Experimental.
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
	// Experimental.
	Type AssetType `field:"optional" json:"type" yaml:"type"`
}

