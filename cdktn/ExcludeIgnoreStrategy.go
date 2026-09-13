// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/open-constructs/cdk-terrain-go/cdktn/jsii"
)

// The default ignore strategy: exact paths, `*.ext` suffixes, and directories (with everything inside them).
// Experimental.
type ExcludeIgnoreStrategy interface {
	IIgnoreStrategy
	// A value identifying this strategy's exclusion behavior, suitable for folding into a cache key.
	//
	// Two strategies that return the same
	// `cacheKey` must exclude the same paths.
	//
	// Callers such as `AssetStaging`'s result cache key on a JSON-serializable
	// representation of their inputs, which a strategy instance is not. Omit
	// this when the strategy's behavior can't be summarized this way; the
	// caller then has to treat every call as uncacheable.
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
	// Experimental.
	PruneExcludedDirectories() *bool
	// Whether the given entry should be excluded.
	// Experimental.
	Ignores(query *IgnoreQuery) *bool
}

// The jsii proxy struct for ExcludeIgnoreStrategy
type jsiiProxy_ExcludeIgnoreStrategy struct {
	jsiiProxy_IIgnoreStrategy
}

func (j *jsiiProxy_ExcludeIgnoreStrategy) CacheKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExcludeIgnoreStrategy) PruneExcludedDirectories() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"pruneExcludedDirectories",
		&returns,
	)
	return returns
}


// Experimental.
func NewExcludeIgnoreStrategy(exclude *[]*string) ExcludeIgnoreStrategy {
	_init_.Initialize()

	if err := validateNewExcludeIgnoreStrategyParameters(exclude); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExcludeIgnoreStrategy{}

	_jsii_.Create(
		"cdktn.ExcludeIgnoreStrategy",
		[]interface{}{exclude},
		&j,
	)

	return &j
}

// Experimental.
func NewExcludeIgnoreStrategy_Override(e ExcludeIgnoreStrategy, exclude *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktn.ExcludeIgnoreStrategy",
		[]interface{}{exclude},
		e,
	)
}

func (e *jsiiProxy_ExcludeIgnoreStrategy) Ignores(query *IgnoreQuery) *bool {
	if err := e.validateIgnoresParameters(query); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		e,
		"ignores",
		[]interface{}{query},
		&returns,
	)

	return returns
}

