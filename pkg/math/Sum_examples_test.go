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

func ExampleSum_uint8s() {
	out, err := Sum([]uint8{1, 2, 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 6
}

func ExampleSum_ints() {
	out, err := Sum([]int{10, 20})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 30
}

func ExampleSum_floats() {
	out, err := Sum([]float64{1.11, 2.22})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 3.33
}

func ExampleSum_durations() {
	out, err := Sum([]time.Duration{
		time.Hour * 2,
		time.Hour * 1,
		time.Hour * 3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 6h0m0s
}

func ExampleSum_interface() {
	out, err := Sum([]interface{}{10, 2.22})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 12.22
}
