// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn


// Provider-protocol feature families whose usage is gated on the project's declared `targetVersions`.
//
// Generated provider bindings pass a member of
// this enum to `TerraformResource.registerProviderFeatureUsage()` when a
// versioned feature is actually used; the per-product minimum versions live
// in the (internal) `providerFeatureConstraints` map in this module.
// Experimental.
type ProviderFeature string

const (
	// Provider-defined functions (`provider::<ns>::<fn>()` expressions).
	// Experimental.
	ProviderFeature_PROVIDER_FUNCTIONS ProviderFeature = "PROVIDER_FUNCTIONS"
	// Ephemeral resources (`ephemeral` blocks), which are never persisted to state or plan files.
	// Experimental.
	ProviderFeature_EPHEMERAL_RESOURCES ProviderFeature = "EPHEMERAL_RESOURCES"
	// Write-only resource attributes: accepted on create/update but never persisted to state and never returned by the provider.
	// Experimental.
	ProviderFeature_WRITE_ONLY_ATTRIBUTES ProviderFeature = "WRITE_ONLY_ATTRIBUTES"
	// Resource identity data alongside state.
	//
	// Reserved: not yet consumed by
	// any generated binding.
	// Experimental.
	ProviderFeature_RESOURCE_IDENTITY ProviderFeature = "RESOURCE_IDENTITY"
)

