// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package math provides math functions that support varied types, but tries to preserve the original type if possible.  For example, you can sum a slice of ints and floats.
package math

import (
	"github.com/pkg/errors"
)

var (
	ErrEmptyInput = errors.New("empty input")
)
