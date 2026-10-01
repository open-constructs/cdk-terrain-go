// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Transforms a source tree into a built artifact. Runs at synth, before the output is packaged and staged.
//
// This is the extension point for asset bundling: core ships no bundler.
// Docker, esbuild, pip, `go build`, and similar are an open-ended, non
// cloud-specific set, so each lives in its own package and implements this
// one interface — the same way {@link IIgnoreStrategy} lets richer exclusion
// live outside core. Users pass an instance via the consuming construct's
// `bundler` option.
//
// `bundle` runs during the owning construct's `onSynthesize` hook and may
// touch the filesystem. Deferring it there keeps it skippable when the asset's
// stack is not being synthesized, which holds as long as the hash is taken
// over the source rather than the built output.
//
// Bundlers compose through {@link ChainBundler} rather than a hierarchy: a
// bundler declines (see {@link BundleResult.declined}) instead of failing when
// it cannot run, and the chain falls through to the next.
// Experimental.
type IAssetBundler interface {
	// Produce the artifact and return a {@link BundleResult} describing it.
	//
	// Implementations write into `options.outputDir` and never write back to
	// `options.source`. The returned path must exist and match its declared
	// shape, or staging rejects it.
	//
	// A bundler that cannot run in this environment returns
	// `BundleResult.declined()` so a {@link ChainBundler} can fall through to
	// the next; a thrown error is a hard failure, not a decline.
	// Experimental.
	Bundle(options *BundleOptions) BundleResult
	// A value identifying the build, folded into the asset hash.
	//
	// The source tree alone cannot see the build, so swapping a `node:18` base
	// image for `node:20` would otherwise leave identity unchanged. Under
	// `SOURCE` hashing this is the only channel by which the build reaches
	// identity, so it must serialize every input that can move the output, or a
	// changed build silently reuses a stale artifact. {@link BundlerKey} builds
	// one from an ordered set of parts.
	//
	// Mirrors {@link IIgnoreStrategy.cacheKey}: omit it when the build cannot be
	// summarized as a string, and fall back to `extraHash`.
	// Default: - the build does not contribute to the hash.
	//
	// Experimental.
	BundlerKey() *string
	// The name a single-file artifact is staged under (e.g. `archive.zip`).
	//
	// A file-producing bundler (`BundleResult.file`) staged with
	// `AssetType.FILE` would otherwise take the source path's basename, which is
	// a directory name when the source is a directory. Declaring the name here
	// lets the artifact reflect what the bundler produces. It is static
	// configuration, needed at construction before a deferred `SOURCE` build
	// runs, not the file's runtime name.
	//
	// Valid only for a file-producing bundler: setting it with a directory
	// packaging (anything but `AssetType.FILE`) is rejected at construction.
	// Default: - the source path's basename.
	//
	// Experimental.
	OutputFileName() *string
}

// The jsii proxy for IAssetBundler
type jsiiProxy_IAssetBundler struct {
	_ byte // padding
}

func (i *jsiiProxy_IAssetBundler) Bundle(options *BundleOptions) BundleResult {
	if err := i.validateBundleParameters(options); err != nil {
		panic(err)
	}
	var returns BundleResult

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

func (j *jsiiProxy_IAssetBundler) OutputFileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFileName",
		&returns,
	)
	return returns
}

