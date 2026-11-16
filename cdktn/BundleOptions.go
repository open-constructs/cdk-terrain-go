// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// Options handed to an {@link IAssetBundler} when it runs.
//
// A struct rather than positional parameters: adding a field is additive,
// adding a method parameter is not, and `bundle` is called through JSII where
// that distinction is a breaking-change boundary.
// Experimental.
type BundleOptions struct {
	// A scratch directory the bundler may write into, owned and created by the caller.
	//
	// The bundler produces its output here (or in a subdirectory) and
	// returns a {@link BundleResult} pointing at the finished artifact — see
	// {@link IAssetBundler.bundle}.
	// Experimental.
	OutputDir *string `field:"required" json:"outputDir" yaml:"outputDir"`
	// Absolute path to the asset's source file or directory. The bundler reads from here and must not modify it.
	//
	// Exclusions (`exclude` / `ignoreStrategy`) are already applied: when any
	// are configured this points at a filtered copy, not the original tree, so
	// the bundler reads exactly the file set the asset hash was taken over.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Absolute path to the asset's source as given, before exclusions.
	//
	// Differs from `source` only for a directory source with `exclude` or
	// `ignoreStrategy` set. Read inputs from `source`; use this to resolve what
	// lies around the source (node_modules, lockfiles, symlink targets). Files
	// read here that `source` lacks are not in a `SOURCE` hash: cover them with
	// `bundlerKey`/`extraHash`, or use `AssetHashType.OUTPUT`.
	// Default: - same as `source`.
	//
	// Experimental.
	OriginalSource *string `field:"optional" json:"originalSource" yaml:"originalSource"`
}

