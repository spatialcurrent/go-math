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

// Compare compares 2 values and follows a few simple rules to compare integers and floats.
// Supports types uint8, int32, int64, int, float64, and time.Time.
// Returns an error if not comparable.
// Return value of -1 indicates a is less than b.
// Return value of 1 indicates a is greater than b.
// Return value of 0 indicates a is equivalent to b by value.
func Compare(a interface{}, b interface{}) (int, error) {

	switch a := a.(type) {
	case time.Duration:
		switch b := b.(type) {
		case time.Duration:
			if a < b {
				return -1, nil
			}
			if a > b {
				return 1, nil
			}
			return 0, nil
		}
	case time.Time:
		switch b := b.(type) {
		case time.Time:
			if a.Before(b) {
				return -1, nil
			}
			if a.After(b) {
				return 1, nil
			}
			return 0, nil
		}
	case uint8:
		switch b := b.(type) {
		case uint8:
			if a < b {
				return -1, nil
			}
			if a > b {
				return 1, nil
			}
			return 0, nil
		case int32:
			if int32(a) < b {
				return -1, nil
			}
			if int32(a) > b {
				return 1, nil
			}
			return 0, nil
		case int64:
			if int64(a) < b {
				return -1, nil
			}
			if int64(a) > b {
				return 1, nil
			}
			return 0, nil
		case int:
			if int(a) < b {
				return -1, nil
			}
			if int(a) > b {
				return 1, nil
			}
			return 0, nil
		case float64:
			if float64(a) < b {
				return -1, nil
			}
			if float64(a) > b {
				return 1, nil
			}
			return 0, nil
		}
	case int32:
		switch b := b.(type) {
		case uint8:
			if a < int32(b) {
				return -1, nil
			}
			if a > int32(b) {
				return 1, nil
			}
			return 0, nil
		case int32:
			if a < b {
				return -1, nil
			}
			if a > b {
				return 1, nil
			}
			return 0, nil
		case int64:
			if int64(a) < b {
				return -1, nil
			}
			if int64(a) > b {
				return 1, nil
			}
			return 0, nil
		case int:
			if int(a) < b {
				return -1, nil
			}
			if int(a) > b {
				return 1, nil
			}
			return 0, nil
		case float64:
			if float64(a) < b {
				return -1, nil
			}
			if float64(a) > b {
				return 1, nil
			}
			return 0, nil
		}
	case int64:
		switch b := b.(type) {
		case uint8:
			if a < int64(b) {
				return -1, nil
			}
			if a > int64(b) {
				return 1, nil
			}
			return 0, nil
		case int32:
			if a < int64(b) {
				return -1, nil
			}
			if a > int64(b) {
				return 1, nil
			}
			return 0, nil
		case int64:
			if a < b {
				return -1, nil
			}
			if a > b {
				return 1, nil
			}
			return 0, nil
		case int:
			if a < int64(b) {
				return -1, nil
			}
			if a > int64(b) {
				return 1, nil
			}
			return 0, nil
		case float64:
			if float64(a) < b {
				return -1, nil
			}
			if float64(a) > b {
				return 1, nil
			}
			return 0, nil
		}
	case int:
		switch b := b.(type) {
		case uint8:
			if a < int(b) {
				return -1, nil
			}
			if a > int(b) {
				return 1, nil
			}
			return 0, nil
		case int32:
			if a < int(b) {
				return -1, nil
			}
			if a > int(b) {
				return 1, nil
			}
			return 0, nil
		case int64:
			if int64(a) < b {
				return -1, nil
			}
			if int64(a) > b {
				return 1, nil
			}
			return 0, nil
		case int:
			if a < b {
				return -1, nil
			}
			if a > b {
				return 1, nil
			}
			return 0, nil
		case float64:
			if float64(a) < b {
				return -1, nil
			}
			if float64(a) > b {
				return 1, nil
			}
			return 0, nil
		}
	case float64:
		switch b := b.(type) {
		case uint8:
			if a < float64(b) {
				return -1, nil
			}
			if a > float64(b) {
				return 1, nil
			}
			return 0, nil
		case int32:
			if a < float64(b) {
				return -1, nil
			}
			if a > float64(b) {
				return 1, nil
			}
			return 0, nil
		case int64:
			if a < float64(b) {
				return -1, nil
			}
			if a > float64(b) {
				return 1, nil
			}
			return 0, nil
		case int:
			if a < float64(b) {
				return -1, nil
			}
			if a > float64(b) {
				return 1, nil
			}
			return 0, nil
		case float64:
			if a < b {
				return -1, nil
			}
			if a > b {
				return 1, nil
			}
			return 0, nil
		}
	}

	return 0, &ErrInvalidComparison{A: a, B: b}
}
