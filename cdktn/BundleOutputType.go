// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// The shape of a bundler's output, so staging and packaging can treat a single-file artifact (a tarball, a `.zip`) differently from a directory tree without inferring it from the path.
// Experimental.
type BundleOutputType string

const (
	// The output is a directory tree.
	//
	// Packaged like an unbundled source directory. The default, and the only
	// shape that predates archive support.
	// Experimental.
	BundleOutputType_DIRECTORY BundleOutputType = "DIRECTORY"
	// The output is a single file the bundler already produced in its final form.
	//
	// A tarball or a deterministic `.zip`. Staged verbatim rather than
	// re-archived, so `AssetType.FILE` no longer has to reject a bundler.
	// Experimental.
	BundleOutputType_FILE BundleOutputType = "FILE"
)

