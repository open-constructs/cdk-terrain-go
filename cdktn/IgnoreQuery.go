// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// A single entry presented to an {@link IIgnoreStrategy}.
//
// A struct rather than a positional `relativePath` string so the shape can
// evolve. `isDirectory` is carried from the outset because `.gitignore` /
// `.dockerignore` semantics turn on it: a `foo/` pattern matches a
// directory but not a file named `foo`, and a matcher cannot recover
// that distinction from the path text alone.
// Experimental.
type IgnoreQuery struct {
	// Whether this entry is a directory rather than a file or symlink.
	//
	// The tree walkers set this from `lstat`, so a strategy can honor
	// directory-only patterns without stat-ing the path itself (which it has
	// no root to resolve against). Excluding a directory also excludes
	// everything below it, since the walkers stop descending once a directory
	// is excluded.
	// Experimental.
	IsDirectory *bool `field:"required" json:"isDirectory" yaml:"isDirectory"`
	// `/`-separated path relative to the asset root.
	// Experimental.
	RelativePath *string `field:"required" json:"relativePath" yaml:"relativePath"`
}

