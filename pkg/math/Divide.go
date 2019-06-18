// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package math

import (
	//stdmath "math"
	"time"
)

// Divide divides a by b while following a few simple rules to maximize precision.  Rather than truncating integers, this function converts the input to floats if a is not divisible by b.
//
//	- If integer a is divisible by integer b, then returns an integer.
//	- If integer a is not divisible by integer b, then converts both numbers to floats and returns the remainder.
// Supports types uint8, int32, int64, int, float64, and time.Duration.
// Catches divide by zero panics and returns as errors.
// Returns an error if a cannot be divided by b.
func Divide(a interface{}, b interface{}) (out interface{}, err error) {

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
	case time.Duration:
		switch b := b.(type) {
		case uint8:
			return time.Duration(int64(float64(int64(a)) / float64(b))), nil
		case int32:
			return time.Duration(int64(float64(int64(a)) / float64(b))), nil
		case int64:
			return time.Duration(int64(float64(int64(a)) / float64(b))), nil
		case int:
			return time.Duration(int64(float64(int64(a)) / float64(b))), nil
		case float64:
			return time.Duration(int64(float64(int64(a)) / float64(b))), nil
		case time.Duration:
			return float64(int64(a)) / float64(int64(b)), nil
		}
	case uint8:
		switch b := b.(type) {
		case uint8:
			if a%b == 0 {
				return a / b, nil
			}
			return float64(a) / float64(b), nil
		case int32:
			if int32(a)%b == 0 {
				return int32(a) / b, nil
			}
			return float64(a) / float64(b), nil
		case int64:
			if int64(a)%b == 0 {
				return int64(a) / b, nil
			}
			return float64(a) / float64(b), nil
		case int:
			if int(a)%b == 0 {
				return int(a) / b, nil
			}
			return float64(a) / float64(b), nil
		case float64:
			return float64(a) / b, nil
		}
	case int32:
		switch b := b.(type) {
		case uint8:
			if a%int32(b) == 0 {
				return a / int32(b), nil
			}
			return float64(a) / float64(b), nil
		case int32:
			if a%b == 0 {
				return a / b, nil
			}
			return float64(a) / float64(b), nil
		case int64:
			if int64(a)%b == 0 {
				return int64(a) / b, nil
			}
			return float64(a) / float64(b), nil
		case int:
			if int(a)%b == 0 {
				return int(a) / b, nil
			}
			return float64(a) / float64(b), nil
		case float64:
			return float64(a) / b, nil
		}
	case int64:
		switch b := b.(type) {
		case uint8:
			if a%int64(b) == 0 {
				return a / int64(b), nil
			}
			return float64(a) / float64(b), nil
		case int32:
			if a%int64(b) == 0 {
				return a / int64(b), nil
			}
			return float64(a) / float64(b), nil
		case int64:
			if a%b == 0 {
				return a / b, nil
			}
			return float64(a) / float64(b), nil
		case int:
			if a%int64(b) == 0 {
				return a / int64(b), nil
			}
			return float64(a) / float64(b), nil
		case float64:
			return float64(a) / b, nil
		}
	case int:
		switch b := b.(type) {
		case uint8:
			if a%int(b) == 0 {
				return a / int(b), nil
			}
			return float64(a) / float64(b), nil
		case int32:
			if a%int(b) == 0 {
				return a / int(b), nil
			}
			return float64(a) / float64(b), nil
		case int64:
			if int64(a)%b == 0 {
				return int64(a) / b, nil
			}
			return float64(a) / float64(b), nil
		case int:
			if a%b == 0 {
				return a / b, nil
			}
			return float64(a) / float64(b), nil
		case float64:
			return float64(a) / b, nil
		}
	case float64:
		switch b := b.(type) {
		case uint8:
			return a / float64(b), nil
		case int32:
			return a / float64(b), nil
		case int64:
			return a / float64(b), nil
		case int:
			return a / float64(b), nil
		case float64:
			return a / b, nil
		}
	}

	return 0, &ErrInvalidDivision{A: a, B: b}
}
