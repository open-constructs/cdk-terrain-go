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
	// Reject a symlink whose target resolves outside `source` and is not itself absolute, rather than packaging a link that dangles once the artifact is extracted elsewhere.
	//
	// Only `ZipPackaging` acts on this — a
	// verbatim directory copy stays on the same machine, so such a link can
	// still be followed there even though it is not portable.
	// Default: false.
	//
	// Experimental.
	RejectEscapingSymlinks *bool `field:"optional" json:"rejectEscapingSymlinks" yaml:"rejectEscapingSymlinks"`
}

