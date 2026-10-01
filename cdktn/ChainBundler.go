// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/open-constructs/cdk-terrain-go/cdktn/jsii"
)

// Composes bundlers into a "try each in order until one runs" chain.
//
// This is how a local bundler and a Docker bundler compose without being
// rewritten as one: each leg declines (`BundleResult.declined()`) when it
// cannot run here, and the chain moves to the next. The first non-declining
// result wins; if all decline, `bundle` throws, since staging has nothing to
// fall back to. Each leg builds into its own output directory, so a leg that
// writes before declining cannot leak partial files into the leg that wins.
//
// The chain's `bundlerKey` folds in *every* leg's key, so identity is the same
// regardless of which leg ends up running — a build that could have gone local
// or Docker is one asset, not two.
//
// This makes the legs interchangeable only if they produce equivalent output.
// Under `SOURCE` hashing they must: a machine with a host tool and one without
// run different legs, and non-equivalent legs would stage different bytes under
// the same hash. Use `OUTPUT` hashing when legs may diverge, so identity tracks
// the artifact each leg actually produced.
// Experimental.
type ChainBundler interface {
	IAssetBundler
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
	// Experimental.
	OutputFileName() *string
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
}

// The jsii proxy struct for ChainBundler
type jsiiProxy_ChainBundler struct {
	jsiiProxy_IAssetBundler
}

func (j *jsiiProxy_ChainBundler) BundlerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundlerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChainBundler) OutputFileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFileName",
		&returns,
	)
	return returns
}


// Chain bundlers in the order given;
//
// earlier bundlers are preferred.
// Experimental.
func ChainBundler_Of(bundlers ...IAssetBundler) ChainBundler {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range bundlers {
		args = append(args, a)
	}

	var returns ChainBundler

	_jsii_.StaticInvoke(
		"cdktn.ChainBundler",
		"of",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChainBundler) Bundle(options *BundleOptions) BundleResult {
	if err := c.validateBundleParameters(options); err != nil {
		panic(err)
	}
	var returns BundleResult

	_jsii_.Invoke(
		c,
		"bundle",
		[]interface{}{options},
		&returns,
	)

	return returns
}

