// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cdktf


// Experimental.
type TerraformResourceMoveByTarget struct {
	// Experimental.
	MoveTarget *string `field:"required" json:"moveTarget" yaml:"moveTarget"`
	// Experimental.
	Index interface{} `field:"optional" json:"index" yaml:"index"`
}

