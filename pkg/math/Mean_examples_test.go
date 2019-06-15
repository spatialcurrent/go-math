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

func ExampleMean_uint8s() {
	out, err := Mean([]uint8{1, 2, 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2
}

func ExampleMean_ints() {
	out, err := Mean([]int{10, 20})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 15
}

func ExampleMean_floats() {
	out, err := Mean([]float64{1.11, 2.22})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1.665
}

func ExampleMean_durations() {
	out, err := Mean([]time.Duration{
		time.Hour * 2,
		time.Hour * 1,
		time.Hour * 3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2h0m0s
}

func ExampleMean_interface() {
	out, err := Mean([]interface{}{10, 2.22})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 6.11
}
