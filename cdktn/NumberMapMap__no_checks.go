// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cdktn

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NumberMapMap) validateLookupParameters(key *string) error {
	return nil
}

func (n *jsiiProxy_NumberMapMap) validateResolveParameters(context IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NumberMapMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NumberMapMap) validateSetTerraformResourceParameters(val IInterpolatingParent) error {
	return nil
}

func validateNewNumberMapMapParameters(terraformResource IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

