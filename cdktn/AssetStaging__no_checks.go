// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cdktn

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetStaging) validateStageParameters(targetPath *string) error {
	return nil
}

func validateAssetStaging_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAssetStagingParameters(scope constructs.Construct, id *string, props *AssetStagingOptions) error {
	return nil
}

