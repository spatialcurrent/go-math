// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package math

import (
	"fmt"
	"time"
)

func ExampleMin_uint8s() {
	out, err := Min([]uint8{1, 2, 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1
}

func ExampleMin_ints() {
	out, err := Min([]int{10, 20})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 10
}

func ExampleMin_floats() {
	out, err := Min([]float64{1.11, 2.22})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1.11
}

func ExampleMin_durations() {
	out, err := Min([]time.Duration{
		time.Hour * 2,
		time.Hour * 3,
		time.Hour * 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1h0m0s
}

func ExampleMin_times() {
	out, err := Min([]time.Time{
		time.Now(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func ExampleMin_interface() {
	out, err := Min([]interface{}{10, 2.22})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2.22
}
