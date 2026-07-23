// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package cdktn

import (
	"fmt"
)

func validateTerraformProviderFunction_InvokeParameters(providerLocalName *string, functionName *string, args *[]interface{}) error {
	if providerLocalName == nil {
		return fmt.Errorf("parameter providerLocalName is required, but nil was provided")
	}

	if functionName == nil {
		return fmt.Errorf("parameter functionName is required, but nil was provided")
	}

	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

