// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package cdktn

import (
	"fmt"
)

func (b *jsiiProxy_BundlerKey) validateWithEnvParameters(entries *map[string]*string) error {
	if entries == nil {
		return fmt.Errorf("parameter entries is required, but nil was provided")
	}

	return nil
}

