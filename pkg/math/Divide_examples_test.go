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

func ExampleDivide_uInt8s() {
	out, err := Divide(uint8(4), uint8(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2
}

func ExampleDivide_ints() {
	out, err := Divide(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 0.5
}

func ExampleDivide_int32s() {
	out, err := Divide(int32(1), int32(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 0.5
}

func ExampleDivide_int64s() {
	out, err := Divide(int64(1), int64(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 0.5
}

func ExampleDivide_floats() {
	out, err := Divide(1.11, 2.22)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 0.5
}

func ExampleDivide_durations() {
	out, err := Divide(time.Hour*3, time.Hour*2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1.5
}

func ExampleDivide_durationInt() {
	out, err := Divide(time.Hour*4, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2h0m0s
}
