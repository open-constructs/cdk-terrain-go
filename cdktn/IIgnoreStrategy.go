// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Decides whether a path relative to an asset root is excluded from staging and hashing.
//
// Core ships only the exact-path / `*.ext` / directory matcher used by
// `exclude` today (`ExcludeIgnoreStrategy`). Full glob, `.gitignore`, and
// `.dockerignore` parity can be implemented against this interface without
// core taking on a glob parser.
// Experimental.
type IIgnoreStrategy interface {
	// Whether the given entry should be excluded.
	// Experimental.
	Ignores(query *IgnoreQuery) *bool
	// A value identifying this strategy's exclusion behavior, suitable for folding into a cache key.
	//
	// Two strategies that return the same
	// `cacheKey` must exclude the same paths.
	//
	// Callers such as `AssetStaging`'s result cache key on a JSON-serializable
	// representation of their inputs, which a strategy instance is not. Omit
	// this when the strategy's behavior can't be summarized this way; the
	// caller then has to treat every call as uncacheable.
	// Default: - this strategy cannot be represented in a cache key.
	//
	// Experimental.
	CacheKey() *string
	// Whether excluding a directory also excludes everything beneath it.
	//
	// When `true` (the default), the walkers stop descending as soon as a
	// directory is excluded — cheaper, and correct for a strategy whose
	// patterns never re-include a path below an excluded parent.
	//
	// A strategy with negation patterns must set this to `false`:
	// `.gitignore` / `.dockerignore` allow `node_modules` followed by
	// `!node_modules/keep`, which is only reachable if the walk descends into
	// the excluded `node_modules` and asks about `node_modules/keep`. The
	// excluded directory entry itself is still omitted; only the descent
	// changes. Opting out costs a full walk of excluded subtrees.
	// Default: true.
	//
	// Experimental.
	PruneExcludedDirectories() *bool
}

// The jsii proxy for IIgnoreStrategy
type jsiiProxy_IIgnoreStrategy struct {
	_ byte // padding
}

func (i *jsiiProxy_IIgnoreStrategy) Ignores(query *IgnoreQuery) *bool {
	if err := i.validateIgnoresParameters(query); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		i,
		"ignores",
		[]interface{}{query},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIgnoreStrategy) CacheKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIgnoreStrategy) PruneExcludedDirectories() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"pruneExcludedDirectories",
		&returns,
	)
	return returns
}

