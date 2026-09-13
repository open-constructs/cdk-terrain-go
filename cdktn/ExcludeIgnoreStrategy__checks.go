// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package cdktn

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (e *jsiiProxy_ExcludeIgnoreStrategy) validateIgnoresParameters(query *IgnoreQuery) error {
	if query == nil {
		return fmt.Errorf("parameter query is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(query, func() string { return "parameter query" }); err != nil {
		return err
	}

	return nil
}

func validateNewExcludeIgnoreStrategyParameters(exclude *[]*string) error {
	if exclude == nil {
		return fmt.Errorf("parameter exclude is required, but nil was provided")
	}

	return nil
}

