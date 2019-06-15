// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package math

import (
	stdmath "math"
	"reflect"
	"time"
)

// Mean returns the mean value from the array or slice of ints, floats, or durations.
// For durations, returns the average duration.
// For ints and floats, returns a float64.
func Mean(in interface{}) (interface{}, error) {

	switch in := in.(type) {
	case []uint8:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := 0
		for i := 0; i < len(in); i++ {
			out += int(in[i])
		}
		return float64(out) / float64(len(in)), nil
	case []int:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := 0
		for i := 0; i < len(in); i++ {
			out += int(in[i])
		}
		return float64(out) / float64(len(in)), nil
	case []int32:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := 0
		for i := 0; i < len(in); i++ {
			out += int(in[i])
		}
		return float64(out) / float64(len(in)), nil
	case []int64:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := int64(0)
		for i := 0; i < len(in); i++ {
			out += in[i]
		}
		return float64(out) / float64(len(in)), nil
	case []float64:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := 0.0
		for i := 0; i < len(in); i++ {
			out += in[i]
		}
		return out / float64(len(in)), nil
	case []time.Duration:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := int64(0)
		for i := 0; i < len(in); i++ {
			out += int64(in[i])
		}
		return time.Duration(int64(stdmath.Floor(float64(out) / float64(len(in))))), nil
	}

	v := reflect.ValueOf(in)
	t := v.Type()
	k := t.Kind()

	if k != reflect.Array && k != reflect.Slice {
		return nil, &ErrInvalidKind{Value: reflect.TypeOf(in), Expected: []reflect.Kind{reflect.Array, reflect.Slice}}
	}

	if v.Len() == 0 {
		return nil, ErrEmptyInput
	}

	sum, err := Sum(in)
	if err != nil {
		return nil, err
	}

	switch sum := sum.(type) {
	case uint8:
		return float64(sum) / float64(v.Len()), nil
	case int32:
		return float64(sum) / float64(v.Len()), nil
	case int64:
		return float64(sum) / float64(v.Len()), nil
	case int:
		return float64(sum) / float64(v.Len()), nil
	case float64:
		return sum / float64(v.Len()), nil
	case time.Duration:
		return time.Duration(int64(stdmath.Floor(float64(int64(sum)) / float64(v.Len())))), nil
	}

	return nil, &ErrInvalidKind{Value: reflect.TypeOf(sum), Expected: []reflect.Kind{reflect.Array, reflect.Slice}}
}
