// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cdktn

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AnyMapMap) validateLookupParameters(key *string) error {
	return nil
}

func (a *jsiiProxy_AnyMapMap) validateResolveParameters(context IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AnyMapMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AnyMapMap) validateSetTerraformResourceParameters(val IInterpolatingParent) error {
	return nil
}

func validateNewAnyMapMapParameters(terraformResource IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

