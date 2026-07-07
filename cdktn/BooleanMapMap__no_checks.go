// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cdktn

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BooleanMapMap) validateLookupParameters(key *string) error {
	return nil
}

func (b *jsiiProxy_BooleanMapMap) validateResolveParameters(context IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BooleanMapMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BooleanMapMap) validateSetTerraformResourceParameters(val IInterpolatingParent) error {
	return nil
}

func validateNewBooleanMapMapParameters(terraformResource IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

