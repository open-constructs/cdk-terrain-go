// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// Options for {@link AssetStaging}.
// Experimental.
type AssetStagingOptions struct {
	// Specify a custom hash for this asset.
	//
	// If `assetHashType` is set it must
	// be set to `AssetHashType.CUSTOM`. The value is used verbatim as the asset
	// hash, and because it names the staged asset file it may only contain
	// letters, digits, `_`, `.` and `-`.
	//
	// NOTE: the hash is used in order to identify a specific revision of the asset, and
	// used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to cloud storage, etc. If you chose to customize the hash, you will
	// need to make sure it is updated every time the asset changes, or otherwise it is
	// possible that some deployments will not be invalidated.
	// Default: - based on `assetHashType`.
	//
	// Experimental.
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
	// Specifies the type of hash to calculate for this asset.
	//
	// If `assetHash` is configured, this option must be `undefined` or
	// `AssetHashType.CUSTOM`.
	// Default: - the default is `AssetHashType.SOURCE`, but if `assetHash` is
	// explicitly specified this value defaults to `AssetHashType.CUSTOM`.
	//
	// Experimental.
	AssetHashType AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// How the staged result is produced and shaped.
	//
	// The caller decides this
	// (e.g. from its own `AssetType`) — `AssetStaging` never infers or changes
	// it based on `exclude`/`extraHash`.
	// Experimental.
	Packaging IAssetPackaging `field:"required" json:"packaging" yaml:"packaging"`
	// Absolute path to the source file or directory.
	//
	// Resolving a relative path
	// against `cdktf.json` is the caller's responsibility.
	// Experimental.
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
	// Paths to exclude, relative to `sourcePath`.
	//
	// Cannot be combined with
	// `ignoreStrategy`, which replaces this matcher rather than layering on
	// top of it.
	// Default: - nothing is excluded.
	//
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Extra information to fold into the hash (e.g. build instructions and other inputs).
	// Default: - no extra hash.
	//
	// Experimental.
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Exclusion matching, for callers that need `.gitignore` / `.dockerignore` parity rather than the built-in exact-path / suffix / directory matcher.
	// Default: - `exclude` is used with the built-in matcher.
	//
	// Experimental.
	IgnoreStrategy IIgnoreStrategy `field:"optional" json:"ignoreStrategy" yaml:"ignoreStrategy"`
}

