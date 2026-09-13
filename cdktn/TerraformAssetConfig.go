// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// Experimental.
type TerraformAssetConfig struct {
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Experimental.
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
	// How the `assetHash` is derived.
	//
	// `SOURCE` (the default) hashes the source path. `CUSTOM` uses the
	// `assetHash` value verbatim and requires it to be set. `OUTPUT` is not
	// supported yet — there is no bundling step to produce an output to hash —
	// and throws if requested.
	//
	// If `assetHash` is set, this must be `undefined` or `AssetHashType.CUSTOM`.
	// Default: AssetHashType.SOURCE
	//
	// Experimental.
	AssetHashType AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// Experimental.
	Type AssetType `field:"optional" json:"type" yaml:"type"`
}

