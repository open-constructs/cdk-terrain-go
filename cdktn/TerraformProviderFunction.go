// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/open-constructs/cdk-terrain-go/cdktn/jsii"
)

// Runtime entry point invoked by generated provider function bindings (`providers/<provider>/provider-functions/index.ts`). Generated code only imports the public `cdktn` package root, so this is the single chokepoint through which every provider-defined function call flows.
// Experimental.
type TerraformProviderFunction interface {
}

// The jsii proxy struct for TerraformProviderFunction
type jsiiProxy_TerraformProviderFunction struct {
	_ byte // padding
}

// Invokes a provider-defined function (Terraform's `provider::<providerLocalName>::<functionName>(...)` syntax).
//
// Note: provider-defined functions are evaluated by the provider itself —
// do not call this inside the configuration of the same provider
// (Terraform reports a self-referential cycle).
// Experimental.
func TerraformProviderFunction_Invoke(providerLocalName *string, functionName *string, args *[]interface{}) IResolvable {
	_init_.Initialize()

	if err := validateTerraformProviderFunction_InvokeParameters(providerLocalName, functionName, args); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktn.TerraformProviderFunction",
		"invoke",
		[]interface{}{providerLocalName, functionName, args},
		&returns,
	)

	return returns
}

