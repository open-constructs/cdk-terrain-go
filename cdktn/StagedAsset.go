// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// A staged artifact, ready to hand to an `IAssetPublisher`.
//
// Deliberately narrower than a location: `path` and `isDirectory` are known
// once staging runs, at synth time, before anything is published. Where an
// asset ends up — a bucket name, an object key, a URL — is resolved at
// apply time and belongs on the publisher's own reference type instead.
// Experimental.
type StagedAsset struct {
	// A hash on the content source.
	//
	// This hash is used to uniquely identify this
	// asset throughout the system. If this value doesn't change, the asset will
	// not be rebuilt or republished.
	// Experimental.
	AssetHash *string `field:"required" json:"assetHash" yaml:"assetHash"`
	// Whether the staged artifact is a directory rather than a single file.
	//
	// Publishers branch on this to decide whether they upload one object or
	// sync a tree; see `IAssetPackaging.producesDirectory`.
	// Experimental.
	IsDirectory *bool `field:"required" json:"isDirectory" yaml:"isDirectory"`
	// The path to the staged artifact, relative to the stack directory.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
}

