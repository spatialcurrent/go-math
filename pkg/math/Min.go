// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package math

import (
	"reflect"
	"time"
)

// Min returns the minimum value from the array or slice of durations, floats, ints, and times.
func Min(in interface{}) (interface{}, error) {

	switch in := in.(type) {
	case []uint8:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := in[0]
		for i := 1; i < len(in); i++ {
			if in[i] < out {
				out = in[i]
			}
		}
		return out, nil
	case []int:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := in[0]
		for i := 1; i < len(in); i++ {
			if in[i] < out {
				out = in[i]
			}
		}
		return out, nil
	case []int32:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := in[0]
		for i := 1; i < len(in); i++ {
			if in[i] < out {
				out = in[i]
			}
		}
		return out, nil
	case []int64:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := in[0]
		for i := 1; i < len(in); i++ {
			if in[i] < out {
				out = in[i]
			}
		}
		return out, nil
	case []float64:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := in[0]
		for i := 1; i < len(in); i++ {
			if in[i] < out {
				out = in[i]
			}
		}
		return out, nil
	case []time.Duration:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := in[0]
		for i := 1; i < len(in); i++ {
			if in[i] < out {
				out = in[i]
			}
		}
		return out, nil
	case []time.Time:
		if len(in) == 0 {
			return nil, ErrEmptyInput
		}
		out := in[0]
		for i := 1; i < len(in); i++ {
			if in[i].Before(out) {
				out = in[i]
			}
		}
		return out, nil
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

	out := v.Index(0)
	for i := 1; i < v.Len(); i++ {
		vi := v.Index(i)
		comparison, err := Compare(vi.Interface(), out.Interface())
		if err != nil {
			return nil, err
		}
		if comparison < 0 {
			out = vi
		}
	}
	return out.Interface(), nil
}
