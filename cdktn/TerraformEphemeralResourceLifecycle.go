// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// Lifecycle options supported by Terraform ephemeral blocks.
//
// Unlike managed
// resources, ephemeral resources have no state, so Terraform only supports
// `precondition`/`postcondition` here - `createBeforeDestroy`,
// `preventDestroy`, `ignoreChanges`, and `replaceTriggeredBy` are all
// state-oriented concepts that do not apply.
// Experimental.
type TerraformEphemeralResourceLifecycle struct {
	// Experimental.
	Postcondition *[]*Postcondition `field:"optional" json:"postcondition" yaml:"postcondition"`
	// Experimental.
	Precondition *[]*Precondition `field:"optional" json:"precondition" yaml:"precondition"`
}

