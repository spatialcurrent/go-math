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

func ExampleMax_uint8s() {
	out, err := Max([]uint8{1, 2, 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 3
}

func ExampleMax_ints() {
	out, err := Max([]int{10, 20})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 20
}

func ExampleMax_floats() {
	out, err := Max([]float64{1.11, 2.22})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2.22
}

func ExampleMax_durations() {
	out, err := Max([]time.Duration{
		time.Hour * 2,
		time.Hour * 3,
		time.Hour * 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 3h0m0s
}

func ExampleMax_times() {
	out, err := Max([]time.Time{
		time.Now(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func ExampleMax_interface() {
	out, err := Max([]interface{}{10, 2.22})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 10
}
