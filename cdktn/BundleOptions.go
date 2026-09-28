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
	// returns the directory that holds the finished artifact — see
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
}

