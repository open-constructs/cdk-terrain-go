// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type ITerraformEphemeralResource interface {
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) IResolvable
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(c interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(d *[]*string)
	// Experimental.
	ForEach() ITerraformIterator
	// Experimental.
	SetForEach(f ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *TerraformEphemeralResourceLifecycle
	// Experimental.
	SetLifecycle(l *TerraformEphemeralResourceLifecycle)
	// Experimental.
	Provider() TerraformProvider
	// Experimental.
	SetProvider(p TerraformProvider)
	// Experimental.
	TerraformResourceType() *string
}

// The jsii proxy for ITerraformEphemeralResource
type jsiiProxy_ITerraformEphemeralResource struct {
	_ byte // padding
}

func (i *jsiiProxy_ITerraformEphemeralResource) InterpolationForAttribute(terraformAttribute *string) IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITerraformEphemeralResource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformEphemeralResource)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ITerraformEphemeralResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformEphemeralResource)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ITerraformEphemeralResource) ForEach() ITerraformIterator {
	var returns ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformEphemeralResource)SetForEach(val ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ITerraformEphemeralResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformEphemeralResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformEphemeralResource) Lifecycle() *TerraformEphemeralResourceLifecycle {
	var returns *TerraformEphemeralResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformEphemeralResource)SetLifecycle(val *TerraformEphemeralResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ITerraformEphemeralResource) Provider() TerraformProvider {
	var returns TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformEphemeralResource)SetProvider(val TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ITerraformEphemeralResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

