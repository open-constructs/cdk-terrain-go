// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testingmatchers


// Experimental.
type ToPlanSuccessfullyOptions struct {
	// Whether `init`/`plan` should use the stack's real backend.
	//
	// Disable this for stacks whose
	// backend (e.g. s3, remote) is not reachable in the test environment and whose plan does not
	// depend on existing state.
	// Default: true.
	//
	// Experimental.
	Backend *bool `field:"optional" json:"backend" yaml:"backend"`
}

