// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// How a staged asset is produced and what shape it takes on disk.
//
// Packaging answers two independent questions: how the artifact is produced
// (copy, zip, tar.gz, ...) and whether the result is a directory or a single
// file. A closed enum can only ever answer the second one, so it is an
// interface rather than an enum — custom formats (e.g. `tar.bz2`) need no
// core change.
// Experimental.
type IAssetPackaging interface {
	// Perform the staging transformation, writing the packaged result to `options.target`.
	// Experimental.
	Pack(options *PackOptions)
	// Appended to the staged artifact name, e.g. ".zip", "", ".tar.bz2".
	// Experimental.
	Extension() *string
	// Whether the staged result is a directory rather than a single file.
	//
	// Publishers branch on this to decide whether they upload one object or
	// sync a tree.
	// Experimental.
	ProducesDirectory() *bool
}

// The jsii proxy for IAssetPackaging
type jsiiProxy_IAssetPackaging struct {
	_ byte // padding
}

func (i *jsiiProxy_IAssetPackaging) Pack(options *PackOptions) {
	if err := i.validatePackParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"pack",
		[]interface{}{options},
	)
}

func (j *jsiiProxy_IAssetPackaging) Extension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssetPackaging) ProducesDirectory() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"producesDirectory",
		&returns,
	)
	return returns
}

