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
	// The hash identifies a specific revision of the asset and caches deployment
	// work (packaging, uploading). A custom hash must be updated whenever the
	// asset changes, or some deployments will not be invalidated.
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
	// A bundler that builds the source into an artifact before staging.
	//
	// Under the default `SOURCE` hashing the build is deferred to `stage()` and
	// stays skippable; `OUTPUT` hashing builds eagerly at construction time to
	// hash the artifact, forgoing skippability. The bundler's declared output
	// shape must match the packaging: a `BundleResult.directory` needs a
	// directory-accepting packaging, a `BundleResult.file` a single-file one
	// (`AssetType.FILE`). The mismatch is caught once the build runs.
	// Default: - the source is staged verbatim, with no build step.
	//
	// Experimental.
	Bundler IAssetBundler `field:"optional" json:"bundler" yaml:"bundler"`
	// Identifier used in error messages, so they name the user-facing construct (e.g. the `TerraformAsset`) rather than this internal staging child.
	// Default: - the staging construct's own id.
	//
	// Experimental.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
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

