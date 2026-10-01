// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// Options controlling how an asset's hash is derived.
// Experimental.
type AssetOptions struct {
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
}

