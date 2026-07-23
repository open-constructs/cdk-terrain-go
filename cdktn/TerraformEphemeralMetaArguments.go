// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// Meta-arguments accepted by ephemeral resources.
//
// Terraform ephemeral blocks
// do not support provisioners or connection blocks, unlike managed resources.
// Experimental.
type TerraformEphemeralMetaArguments struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
}

