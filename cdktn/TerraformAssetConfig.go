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
	// `assetHash` value verbatim and requires it to be set. `OUTPUT` also
	// hashes the source path today — there is no bundling step yet, so the
	// "output" of an asset is its source verbatim — but will hash the
	// bundler's output once bundling is introduced.
	//
	// If `assetHash` is set, this must be `undefined` or `AssetHashType.CUSTOM`.
	// Default: AssetHashType.SOURCE
	//
	// Experimental.
	AssetHashType AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// Paths to exclude from the asset, relative to `path`.
	//
	// See
	// `AssetStagingOptions.exclude` for the accepted forms. Both the computed
	// hash and the staged/packed content honor the exclusion.
	// Default: - nothing is excluded.
	//
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Extra information to fold into the hash (e.g. build instructions and other inputs).
	// Default: - no extra hash.
	//
	// Experimental.
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Experimental.
	Type AssetType `field:"optional" json:"type" yaml:"type"`
}

