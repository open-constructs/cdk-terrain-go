// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/open-constructs/cdk-terrain-go/cdktn/jsii"
)

// Computes a content hash of a file or directory without staging it.
//
// Providers that read a local path directly — a Docker build `context`, for
// example — need a content hash to drive `triggers`, but have no use for a
// staged copy of the source. `Asset` and `TerraformAsset` always produce a
// staged copy; this is the identity half without the staging half.
// Experimental.
type AssetHash interface {
}

// The jsii proxy struct for AssetHash
type jsiiProxy_AssetHash struct {
	_ byte // padding
}

// Content hash of a file or directory, without staging it.
//
// This is a hash of the source tree, and it does not depend on how the
// source is later packaged. `hashPath` is always called with `archive`
// unset, so directory records are part of the digest; the hash of a given
// tree is therefore the same whether it is later copied, zipped, or packed
// by a custom `IAssetPackaging` such as `tar.bz2`.
//
// A consequence worth stating: this does not equal
// `TerraformAsset(dir, { type: ARCHIVE }).assetHash` for a directory with
// subdirectories. That asset frames its hash to the emitted ZIP, which has
// no directory entries (see #323), so directory-only changes move it and
// not this. It equals a `FILE` / `DIRECTORY` `TerraformAsset` hash only
// when the `canonicalAssetHashes` flag is enabled, since that flag is what
// puts the asset on this same canonical scheme.
//
// A relative `filePath` is resolved against the directory containing
// `cdktf.json`, the same base `TerraformAsset` uses, so both hash the same
// source regardless of the process working directory. Absolute paths are
// used as-is. Throws if `filePath` is relative and no `cdktf.json` is found
// above the working directory.
// Experimental.
func AssetHash_Of(filePath *string, options *AssetHashOptions) *string {
	_init_.Initialize()

	if err := validateAssetHash_OfParameters(filePath, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktn.AssetHash",
		"of",
		[]interface{}{filePath, options},
		&returns,
	)

	return returns
}

