// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/open-constructs/cdk-terrain-go/cdktn/jsii"
)

// Builds an {@link IAssetBundler.bundlerKey} from an ordered set of parts.
//
// A `bundlerKey` has to serialize everything that can move a build's output;
// done ad hoc, every bundler invents its own delimiter and forgets an input
// differently. This gives the convention one implementation: parts are joined
// with a separator that is escaped where it appears in a value, so distinct
// inputs can never collide into the same key (`["a:b", "c"]` and
// `["a", "b:c"]` stay different).
//
// Example:
//   const key = BundlerKey.of("docker", image, command)
//     .withEnv({ NODE_ENV: nodeEnv })
//     .toString();
//
// Experimental.
type BundlerKey interface {
	// Append parts, returning a new key.
	// Experimental.
	Add(parts ...*string) BundlerKey
	// Render the collected parts to the string passed as `bundlerKey`.
	// Experimental.
	ToString() *string
	// Append `key=value` parts for a record, sorted by key so the result does not depend on property order.
	// Experimental.
	WithEnv(entries *map[string]*string) BundlerKey
}

// The jsii proxy struct for BundlerKey
type jsiiProxy_BundlerKey struct {
	_ byte // padding
}

// Start a key from an ordered list of parts.
//
// Order is significant: it is part of what the key identifies.
// Experimental.
func BundlerKey_Of(parts ...*string) BundlerKey {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range parts {
		args = append(args, a)
	}

	var returns BundlerKey

	_jsii_.StaticInvoke(
		"cdktn.BundlerKey",
		"of",
		args,
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BundlerKey) Add(parts ...*string) BundlerKey {
	args := []interface{}{}
	for _, a := range parts {
		args = append(args, a)
	}

	var returns BundlerKey

	_jsii_.Invoke(
		b,
		"add",
		args,
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BundlerKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BundlerKey) WithEnv(entries *map[string]*string) BundlerKey {
	if err := b.validateWithEnvParameters(entries); err != nil {
		panic(err)
	}
	var returns BundlerKey

	_jsii_.Invoke(
		b,
		"withEnv",
		[]interface{}{entries},
		&returns,
	)

	return returns
}

