// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Transforms a source tree into a built artifact. Runs at synth, before the output is packaged and staged.
//
// This is the extension point for asset bundling: core ships no bundler.
// Docker, esbuild, pip, `go build`, and similar are an open-ended set that
// is not cloud-specific, so each lives in its own package and implements this
// one interface — the same way {@link IIgnoreStrategy} lets richer exclusion
// live outside core without core taking on a glob parser. A third party
// develops a bundler by implementing this interface and publishing it as a
// package; users pass an instance via the consuming construct's `bundler`
// option.
//
// `bundle` runs during the owning construct's `onSynthesize` hook and may
// touch the filesystem. Deferring it there keeps it skippable when the asset's
// stack is not being synthesized, which holds as long as the hash is taken
// over the source rather than the built output.
// Experimental.
type IAssetBundler interface {
	// Produce the artifact and return the directory holding it.
	//
	// Implementations write into `options.outputDir` and return it or a
	// subdirectory, never writing back to `options.source`. The returned
	// directory is then packaged as an unbundled source directory would be.
	// Returning a file, or a path that does not exist, is rejected — the
	// contract is a directory, and packaging always treats the result as one.
	// Experimental.
	Bundle(options *BundleOptions) *string
	// A value identifying the build, folded into the asset hash.
	//
	// The source tree alone cannot see the build, so swapping a `node:18` base
	// image for `node:20` would otherwise leave identity unchanged. A value
	// capturing the build (e.g. `docker:<image>:<command>`) closes that gap.
	//
	// Under `SOURCE` hashing this is the only channel by which the build reaches
	// identity, so it must serialize every input that can move the output —
	// base image, command, tool version, environment, arguments. Anything left
	// out means a changed build silently reuses a stale artifact. {@link * BundlerKey} builds one from an ordered set of parts so the format is not
	// reinvented per bundler.
	//
	// Mirrors {@link IIgnoreStrategy.cacheKey}: omit it when the build cannot be
	// summarized as a string, and fall back to `extraHash`.
	// Default: - the build does not contribute to the hash.
	//
	// Experimental.
	BundlerKey() *string
}

// The jsii proxy for IAssetBundler
type jsiiProxy_IAssetBundler struct {
	_ byte // padding
}

func (i *jsiiProxy_IAssetBundler) Bundle(options *BundleOptions) *string {
	if err := i.validateBundleParameters(options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"bundle",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAssetBundler) BundlerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundlerKey",
		&returns,
	)
	return returns
}

