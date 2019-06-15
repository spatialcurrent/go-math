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

func ExampleMultiply_uInt8s() {
	out, err := Multiply(uint8(1), uint8(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2
}

func ExampleMultiply_ints() {
	out, err := Multiply(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2
}

func ExampleMultiply_int32s() {
	out, err := Multiply(int32(1), int32(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2
}

func ExampleMultiply_int64s() {
	out, err := Multiply(int64(1), int64(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2
}

func ExampleMultiply_floats() {
	out, err := Multiply(0.5, 3.0)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1.5
}

func ExampleMultiply_durationInt() {
	out, err := Multiply(2, time.Hour*1)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2h0m0s
}

func ExampleMultiply_durationFloat() {
	out, err := Multiply(0.5, time.Hour*3)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1h30m0s
}
