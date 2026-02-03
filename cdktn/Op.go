// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/open-constructs/cdk-terrain-go/cdktn/jsii"
)

// This class contains static functions for all arithmetical and logical operators in the Terraform configuration language.
// Experimental.
type Op interface {
}

// The jsii proxy struct for Op
type jsiiProxy_Op struct {
	_ byte // padding
}

// Experimental.
func NewOp() Op {
	_init_.Initialize()

	j := jsiiProxy_Op{}

	_jsii_.Create(
		"cdktn.Op",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewOp_Override(o Op) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktn.Op",
		nil, // no parameters
		o,
	)
}

// Renders left + right.
// Experimental.
func Op_Add(left interface{}, right interface{}) *float64 {
	_init_.Initialize()

	if err := validateOp_AddParameters(left, right); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"add",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders left && right.
// Experimental.
func Op_And(left interface{}, right interface{}) IResolvable {
	_init_.Initialize()

	if err := validateOp_AndParameters(left, right); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"and",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders left / right.
// Experimental.
func Op_Div(left interface{}, right interface{}) *float64 {
	_init_.Initialize()

	if err := validateOp_DivParameters(left, right); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"div",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders left == right.
// Experimental.
func Op_Eq(left interface{}, right interface{}) IResolvable {
	_init_.Initialize()

	if err := validateOp_EqParameters(left, right); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"eq",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders left > right.
// Experimental.
func Op_Gt(left interface{}, right interface{}) IResolvable {
	_init_.Initialize()

	if err := validateOp_GtParameters(left, right); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"gt",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders left >= right.
// Experimental.
func Op_Gte(left interface{}, right interface{}) IResolvable {
	_init_.Initialize()

	if err := validateOp_GteParameters(left, right); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"gte",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders left < right.
// Experimental.
func Op_Lt(left interface{}, right interface{}) IResolvable {
	_init_.Initialize()

	if err := validateOp_LtParameters(left, right); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"lt",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders left <= right.
// Experimental.
func Op_Lte(left interface{}, right interface{}) IResolvable {
	_init_.Initialize()

	if err := validateOp_LteParameters(left, right); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"lte",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders left % right.
// Experimental.
func Op_Mod(left interface{}, right interface{}) *float64 {
	_init_.Initialize()

	if err := validateOp_ModParameters(left, right); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"mod",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders left * right.
// Experimental.
func Op_Mul(left interface{}, right interface{}) *float64 {
	_init_.Initialize()

	if err := validateOp_MulParameters(left, right); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"mul",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders -expression.
// Experimental.
func Op_Negate(expression interface{}) *float64 {
	_init_.Initialize()

	if err := validateOp_NegateParameters(expression); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"negate",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// Renders left != right.
// Experimental.
func Op_Neq(left interface{}, right interface{}) IResolvable {
	_init_.Initialize()

	if err := validateOp_NeqParameters(left, right); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"neq",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders !expression.
// Experimental.
func Op_Not(expression interface{}) IResolvable {
	_init_.Initialize()

	if err := validateOp_NotParameters(expression); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"not",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// Renders left || right.
// Experimental.
func Op_Or(left interface{}, right interface{}) IResolvable {
	_init_.Initialize()

	if err := validateOp_OrParameters(left, right); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"or",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Renders left - right.
// Experimental.
func Op_Sub(left interface{}, right interface{}) *float64 {
	_init_.Initialize()

	if err := validateOp_SubParameters(left, right); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktn.Op",
		"sub",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

