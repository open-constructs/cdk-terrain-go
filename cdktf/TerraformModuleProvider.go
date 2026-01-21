// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cdktf


// Experimental.
type TerraformModuleProvider struct {
	// Experimental.
	ModuleAlias *string `field:"required" json:"moduleAlias" yaml:"moduleAlias"`
	// Experimental.
	Provider TerraformProvider `field:"required" json:"provider" yaml:"provider"`
}

