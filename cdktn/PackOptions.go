// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// Options for {@link IAssetPackaging.pack}.
//
// A struct rather than positional parameters: adding a struct field is
// additive, adding a method parameter is not, and `pack` is called through
// JSII where that distinction is a breaking-change boundary.
// Experimental.
type PackOptions struct {
	// Path to the resolved (already bundled, if applicable) source.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Path the packaged result should be written to.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Entries to omit from the packaged result.
	//
	// Must match the strategy used
	// to hash the same source, or the hash and the artifact describe
	// different sets of files.
	// Default: - nothing is excluded.
	//
	// Experimental.
	IgnoreStrategy IIgnoreStrategy `field:"optional" json:"ignoreStrategy" yaml:"ignoreStrategy"`
}

