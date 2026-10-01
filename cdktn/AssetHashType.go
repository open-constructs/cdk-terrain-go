// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// The type of asset hash.
//
// The hash identifies a specific revision of the asset and caches deployment
// work such as packaging and uploading.
// Experimental.
type AssetHashType string

const (
	// Based on the content of the source path.
	//
	// Use `SOURCE` to track changes to the source files directly.
	// Experimental.
	AssetHashType_SOURCE AssetHashType = "SOURCE"
	// Based on the content of the bundling output.
	//
	// Use `OUTPUT` when the source is a top-level folder holding code and/or
	// dependencies not directly linked to the asset.
	// Experimental.
	AssetHashType_OUTPUT AssetHashType = "OUTPUT"
	// Use a custom hash.
	// Experimental.
	AssetHashType_CUSTOM AssetHashType = "CUSTOM"
)

