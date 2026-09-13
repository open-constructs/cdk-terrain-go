// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cdktn

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExcludeIgnoreStrategy) validateIgnoresParameters(query *IgnoreQuery) error {
	return nil
}

func validateNewExcludeIgnoreStrategyParameters(exclude *[]*string) error {
	return nil
}

