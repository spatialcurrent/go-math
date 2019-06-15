// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package math

import (
	stdmath "math"
	"time"
)

// Multiply multiplies 2 values together while following a few simple rules to convert integers and floats.
// Supports types uint8, int32, int64, int, float64, and time.Duration.
// Returns an error if the values cannot be multiplied.
func Multiply(a interface{}, b interface{}) (interface{}, error) {

	switch a := a.(type) {
	case time.Duration:
		switch b := b.(type) {
		case uint8:
			return time.Duration(int64(a) * int64(b)), nil
		case int32:
			return time.Duration(int64(a) * int64(b)), nil
		case int64:
			return time.Duration(int64(a) * b), nil
		case int:
			return time.Duration(int64(a) * int64(b)), nil
		case float64:
			return time.Duration(int64(stdmath.Floor(float64(int64(a)) * b))), nil
		}
	case uint8:
		switch b := b.(type) {
		case uint8:
			return a * b, nil
		case int32:
			return int32(a) * b, nil
		case int64:
			return int64(a) * b, nil
		case int:
			return int(a) * b, nil
		case float64:
			return float64(a) * b, nil
		case time.Duration:
			return time.Duration(int64(a) * int64(b)), nil
		}
	case int32:
		switch b := b.(type) {
		case uint8:
			return a * int32(b), nil
		case int32:
			return a * b, nil
		case int64:
			return int64(a) * b, nil
		case int:
			return int(a) * b, nil
		case float64:
			return float64(a) * b, nil
		case time.Duration:
			return time.Duration(int64(a) * int64(b)), nil
		}
	case int64:
		switch b := b.(type) {
		case uint8:
			return a * int64(b), nil
		case int32:
			return a * int64(b), nil
		case int64:
			return a * b, nil
		case int:
			return a * int64(b), nil
		case float64:
			return float64(a) * b, nil
		case time.Duration:
			return time.Duration(a * int64(b)), nil
		}
	case int:
		switch b := b.(type) {
		case uint8:
			return a * int(b), nil
		case int32:
			return a * int(b), nil
		case int64:
			return int64(a) * b, nil
		case int:
			return a * b, nil
		case float64:
			return float64(a) * b, nil
		case time.Duration:
			return time.Duration(int64(a) * int64(b)), nil
		}
	case float64:
		switch b := b.(type) {
		case uint8:
			return a * float64(b), nil
		case int32:
			return a * float64(b), nil
		case int64:
			return a * float64(b), nil
		case int:
			return a * float64(b), nil
		case float64:
			return a * b, nil
		case time.Duration:
			return time.Duration(int64(stdmath.Floor(a * float64(int64(b))))), nil
		}
	}

	return 0, &ErrInvalidMultiplication{A: a, B: b}
}
