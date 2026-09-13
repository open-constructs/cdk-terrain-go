// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/open-constructs/cdk-terrain-go/cdktn/jsii"
)

// Built-in packaging strategies.
//
// Custom formats implement `IAssetPackaging`
// directly rather than extending this class.
// Experimental.
type AssetPackaging interface {
}

// The jsii proxy struct for AssetPackaging
type jsiiProxy_AssetPackaging struct {
	_ byte // padding
}

func AssetPackaging_DIRECTORY() IAssetPackaging {
	_init_.Initialize()
	var returns IAssetPackaging
	_jsii_.StaticGet(
		"cdktn.AssetPackaging",
		"DIRECTORY",
		&returns,
	)
	return returns
}

func AssetPackaging_FILE() IAssetPackaging {
	_init_.Initialize()
	var returns IAssetPackaging
	_jsii_.StaticGet(
		"cdktn.AssetPackaging",
		"FILE",
		&returns,
	)
	return returns
}

func AssetPackaging_ZIP() IAssetPackaging {
	_init_.Initialize()
	var returns IAssetPackaging
	_jsii_.StaticGet(
		"cdktn.AssetPackaging",
		"ZIP",
		&returns,
	)
	return returns
}

