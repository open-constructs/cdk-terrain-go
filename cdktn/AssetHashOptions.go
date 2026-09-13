// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// Options for {@link AssetHash.of}.
// Experimental.
type AssetHashOptions struct {
	// Paths to exclude, relative to the hashed path.
	//
	// Cannot be combined with
	// `ignoreStrategy`, which replaces this matcher rather than layering on
	// top of it.
	// Default: - nothing is excluded.
	//
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Extra information to fold into the hash.
	// Default: - no extra data.
	//
	// Experimental.
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Exclusion matching, for callers that need `.gitignore` / `.dockerignore` parity rather than the built-in exact-path / suffix / directory matcher.
	// Default: - `exclude` is used with the built-in matcher.
	//
	// Experimental.
	IgnoreStrategy IIgnoreStrategy `field:"optional" json:"ignoreStrategy" yaml:"ignoreStrategy"`
}

