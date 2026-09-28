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
	// `assetHash` value verbatim and requires it to be set. `OUTPUT` hashes the
	// source too, unless a `bundler` is set — then it hashes the bundler's
	// built output, which forces an eager build (see `bundler`).
	//
	// If `assetHash` is set, this must be `undefined` or `AssetHashType.CUSTOM`.
	// Default: AssetHashType.SOURCE
	//
	// Experimental.
	AssetHashType AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// A bundler that builds the source into an artifact before staging.
	//
	// Core ships no bundler; implement `IAssetBundler` or use one from a bundler
	// package. Under the default `SOURCE` hashing the build is deferred to synth
	// and stays skippable; `OUTPUT` hashing builds eagerly to hash the artifact.
	// `AssetType.FILE` is rejected, since bundler output is always a directory.
	// See `AssetStagingOptions.bundler`.
	// Default: - the source is staged verbatim, with no build step.
	//
	// Experimental.
	Bundler IAssetBundler `field:"optional" json:"bundler" yaml:"bundler"`
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

