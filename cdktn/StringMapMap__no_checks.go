// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cdktn

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StringMapMap) validateLookupParameters(key *string) error {
	return nil
}

func (s *jsiiProxy_StringMapMap) validateResolveParameters(context IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StringMapMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StringMapMap) validateSetTerraformResourceParameters(val IInterpolatingParent) error {
	return nil
}

func validateNewStringMapMapParameters(terraformResource IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

