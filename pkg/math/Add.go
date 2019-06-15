// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package math

import (
	"time"
)

// Add adds 2 values together while following a few simple rules to convert integers and floats.
// Supports types uint8, int32, int64, int, float64, and time.Duration.
// Returns an error if the values cannot be added.
func Add(a interface{}, b interface{}) (interface{}, error) {

	switch a := a.(type) {
	case time.Duration:
		switch b := b.(type) {
		case time.Duration:
			return time.Duration(a + b), nil
		case time.Time:
			return b.Add(a), nil
		}
	case time.Time:
		switch b := b.(type) {
		case time.Duration:
			return a.Add(b), nil
		}
	case uint8:
		switch b := b.(type) {
		case uint8:
			return a + b, nil
		case int32:
			return int32(a) + b, nil
		case int64:
			return int64(a) + b, nil
		case int:
			return int(a) + b, nil
		case float64:
			return float64(a) + b, nil
		}
	case int32:
		switch b := b.(type) {
		case uint8:
			return a + int32(b), nil
		case int32:
			return a + b, nil
		case int64:
			return int64(a) + b, nil
		case int:
			return int(a) + b, nil
		case float64:
			return float64(a) + b, nil
		}
	case int64:
		switch b := b.(type) {
		case uint8:
			return a + int64(b), nil
		case int32:
			return a + int64(b), nil
		case int64:
			return a + b, nil
		case int:
			return a + int64(b), nil
		case float64:
			return float64(a) + b, nil
		}
	case int:
		switch b := b.(type) {
		case uint8:
			return a + int(b), nil
		case int32:
			return a + int(b), nil
		case int64:
			return int64(a) + b, nil
		case int:
			return a + b, nil
		case float64:
			return float64(a) + b, nil
		}
	case float64:
		switch b := b.(type) {
		case uint8:
			return a + float64(b), nil
		case int32:
			return a + float64(b), nil
		case int64:
			return a + float64(b), nil
		case int:
			return a + float64(b), nil
		case float64:
			return a + b, nil
		}
	}

	return 0, &ErrInvalidAddition{A: a, B: b}
}
