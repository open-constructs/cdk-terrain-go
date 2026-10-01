// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/open-constructs/cdk-terrain-go/cdktn/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn/internal"
)

// Resolves an asset's identity and stages its content to disk.
//
// The hash is available immediately after construction; the filesystem write
// is deferred to `stage()`, the one window where it is safe. Hashing is
// cached per synth (see {@link hashCachesByRoot}) so an asset referenced from
// several places is walked once. Without a bundler, `SOURCE` and `OUTPUT` both
// hash the source; with one they diverge, as `AssetStagingOptions.bundler`
// describes.
// Experimental.
type AssetStaging interface {
	constructs.Construct
	IAsset
	// A hash of this asset, available at construction time.
	//
	// Being a plain string, it can be used in construct IDs to force a new
	// resource when the content hash changes.
	// Experimental.
	AssetHash() *string
	// Experimental.
	IsDirectory() *bool
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Packaging() IAssetPackaging
	// Stage the asset's content to `targetPath`.
	//
	// Called from the owning construct's `onSynthesize` hook. Without a bundler
	// the source is packaged verbatim; with one, the bundler's output is, built
	// eagerly for `OUTPUT` hashing or deferred to here for `SOURCE`.
	// Experimental.
	Stage(targetPath *string)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for AssetStaging
type jsiiProxy_AssetStaging struct {
	internal.Type__constructsConstruct
	jsiiProxy_IAsset
}

func (j *jsiiProxy_AssetStaging) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) IsDirectory() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) Packaging() IAssetPackaging {
	var returns IAssetPackaging
	_jsii_.Get(
		j,
		"packaging",
		&returns,
	)
	return returns
}


// Experimental.
func NewAssetStaging(scope constructs.Construct, id *string, props *AssetStagingOptions) AssetStaging {
	_init_.Initialize()

	if err := validateNewAssetStagingParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetStaging{}

	_jsii_.Create(
		"cdktn.AssetStaging",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetStaging_Override(a AssetStaging, scope constructs.Construct, id *string, props *AssetStagingOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktn.AssetStaging",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func AssetStaging_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAssetStaging_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdktn.AssetStaging",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetStaging) Stage(targetPath *string) {
	if err := a.validateStageParameters(targetPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"stage",
		[]interface{}{targetPath},
	)
}

func (a *jsiiProxy_AssetStaging) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetStaging) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

