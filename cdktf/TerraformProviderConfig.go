// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cdktf


// Experimental.
type TerraformProviderConfig struct {
	// Experimental.
	TerraformResourceType *string `field:"required" json:"terraformResourceType" yaml:"terraformResourceType"`
	// Experimental.
	TerraformGeneratorMetadata *TerraformProviderGeneratorMetadata `field:"optional" json:"terraformGeneratorMetadata" yaml:"terraformGeneratorMetadata"`
	// Experimental.
	TerraformProviderSource *string `field:"optional" json:"terraformProviderSource" yaml:"terraformProviderSource"`
}

