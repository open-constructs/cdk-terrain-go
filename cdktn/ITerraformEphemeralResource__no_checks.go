// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cdktn

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITerraformEphemeralResource) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_ITerraformEphemeralResource) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ITerraformEphemeralResource) validateSetLifecycleParameters(val *TerraformEphemeralResourceLifecycle) error {
	return nil
}

