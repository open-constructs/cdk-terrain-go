// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/open-constructs/cdk-terrain-go/cdktn/jsii"
)

// What a bundler produced, returned from {@link IAssetBundler.bundle}.
//
// Carries the artifact path and its shape, or a declined state signalling the
// caller should fall back. See {@link declined} for the decline protocol.
// Experimental.
type BundleResult interface {
	// Whether the bundler declined to run, signalling the caller to fall back.
	// Experimental.
	IsDeclined() *bool
	// The artifact's shape, or undefined when {@link isDeclined}.
	// Experimental.
	OutputType() BundleOutputType
	// The artifact path, or undefined when {@link isDeclined}.
	// Experimental.
	Path() *string
}

// The jsii proxy struct for BundleResult
type jsiiProxy_BundleResult struct {
	_ byte // padding
}

func (j *jsiiProxy_BundleResult) IsDeclined() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isDeclined",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BundleResult) OutputType() BundleOutputType {
	var returns BundleOutputType
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BundleResult) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


// The bundler declined to run here; the caller should fall back.
//
// Distinct from a thrown error, which is a hard failure.
// Experimental.
func BundleResult_Declined() BundleResult {
	_init_.Initialize()

	var returns BundleResult

	_jsii_.StaticInvoke(
		"cdktn.BundleResult",
		"declined",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A directory-tree artifact at `path`.
// Experimental.
func BundleResult_Directory(path *string) BundleResult {
	_init_.Initialize()

	if err := validateBundleResult_DirectoryParameters(path); err != nil {
		panic(err)
	}
	var returns BundleResult

	_jsii_.StaticInvoke(
		"cdktn.BundleResult",
		"directory",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// A single-file artifact at `path` (a tarball, a `.zip`), staged verbatim.
// Experimental.
func BundleResult_File(path *string) BundleResult {
	_init_.Initialize()

	if err := validateBundleResult_FileParameters(path); err != nil {
		panic(err)
	}
	var returns BundleResult

	_jsii_.StaticInvoke(
		"cdktn.BundleResult",
		"file",
		[]interface{}{path},
		&returns,
	)

	return returns
}

