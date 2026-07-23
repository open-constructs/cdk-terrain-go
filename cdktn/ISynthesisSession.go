// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a single session of synthesis.
//
// Passed into `TerraformStack.onSynthesize()` methods.
// originally from aws/constructs lib v3.3.126 (synth functionality was removed in constructs v10)
// Experimental.
type ISynthesisSession interface {
	// Experimental.
	Manifest() Manifest
	// The output directory for this synthesis session.
	// Experimental.
	Outdir() *string
	// Experimental.
	SkipValidation() *bool
	// Whether every stack that will be synthesized in this session already had `prepareStack()` run against it, for ALL of those stacks, before any of their synthesizers are invoked.
	//
	// `App.synth()` sets this to `true`: it calls `prepareStack()` on every
	// stack up front, so resolve-discovered provider-feature usage (see
	// `TerraformElement._registerResolveDiscoveredProviderFeatureUsage`) and
	// Terraform-function usage are visible across ALL of an App's stacks
	// before ANY of them run validations - a cross-stack reference in stack B
	// that only surfaces usage while resolving stack A must not make stack
	// A's validation miss it.
	//
	// Left unset (or `false`) by sessions built outside `App.synth()` (a
	// `StackSynthesizer` driven directly). `StackSynthesizer.synthesize()`
	// then runs `prepareStack()` itself before validating, so those entry
	// points still discover current-pass usage instead of silently skipping
	// the validations that depend on it - at the cost of only ever seeing
	// that one stack's own usage, since no sibling-stack preparation happened
	// first.
	// Experimental.
	StacksPrepared() *bool
}

// The jsii proxy for ISynthesisSession
type jsiiProxy_ISynthesisSession struct {
	_ byte // padding
}

func (j *jsiiProxy_ISynthesisSession) Manifest() Manifest {
	var returns Manifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession) SkipValidation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession) StacksPrepared() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"stacksPrepared",
		&returns,
	)
	return returns
}

