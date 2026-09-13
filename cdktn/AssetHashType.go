// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// The type of asset hash.
//
// NOTE: the hash is used in order to identify a specific revision of the asset, and
// used for optimizing and caching deployment activities related to this asset such as
// packaging, uploading to cloud storage, etc.
// Experimental.
type AssetHashType string

const (
	// Based on the content of the source path.
	//
	// Use `SOURCE` when the content of the asset changes frequently or when
	// you want to track changes to the source files directly.
	// Experimental.
	AssetHashType_SOURCE AssetHashType = "SOURCE"
	// Based on the content of the bundling output.
	//
	// Use `OUTPUT` when the source of the asset is a top level folder containing
	// code and/or dependencies that are not directly linked to the asset.
	// Experimental.
	AssetHashType_OUTPUT AssetHashType = "OUTPUT"
	// Use a custom hash.
	// Experimental.
	AssetHashType_CUSTOM AssetHashType = "CUSTOM"
)

