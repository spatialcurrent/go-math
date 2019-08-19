// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package math

import (
	stdmath "math"
)

// Pow raises a to the power of b (a^b).
// If a and b are both (unsigned-)integers, then returns an int.  Otherwise, returns a float64.
// Supports types uint8, int32, int64, int, and float64.
func Pow(a interface{}, b interface{}) (out interface{}, err error) {

	// Catch and recover from runtime error, e.g., divide by zero.
	defer func() {
		if r := recover(); r != nil {
			if re, ok := r.(error); ok {
				out = nil
				err = re
			} else {
				panic(r)
			}
		}
	}()

	switch a := a.(type) {
	case uint8:
		switch b := b.(type) {
		case uint8:
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int32:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int64:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case float64:
			return stdmath.Pow(float64(a), b), nil
		}
	case int32:
		switch b := b.(type) {
		case uint8:
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int32:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int64:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case float64:
			return stdmath.Pow(float64(a), b), nil
		}
	case int64:
		switch b := b.(type) {
		case uint8:
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int32:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int64:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case float64:
			return stdmath.Pow(float64(a), b), nil
		}
	case int:
		switch b := b.(type) {
		case uint8:
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int32:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int64:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case int:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return int(stdmath.Pow(float64(a), float64(b))), nil
		case float64:
			return stdmath.Pow(float64(a), b), nil
		}
	case float64:
		switch b := b.(type) {
		case uint8:
			return stdmath.Pow(a, float64(b)), nil
		case int32:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return stdmath.Pow(a, float64(b)), nil
		case int64:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return stdmath.Pow(a, float64(b)), nil
		case int:
			if b < 0 {
				return stdmath.Pow(float64(a), float64(b)), nil
			}
			return stdmath.Pow(a, float64(b)), nil
		case float64:
			return stdmath.Pow(a, b), nil
		}
	}

	return 0, &ErrInvalidPower{A: a, B: b}
}
