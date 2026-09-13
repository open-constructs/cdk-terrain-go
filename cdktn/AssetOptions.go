// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// Asset hash options.
// Experimental.
type AssetOptions struct {
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
}

