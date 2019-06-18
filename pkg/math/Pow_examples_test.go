// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package math

import (
	"fmt"
)

func ExamplePow_uInt8s() {
	out, err := Pow(uint8(4), uint8(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 16
}

func ExamplePow_ints() {
	out, err := Pow(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1
}

func ExamplePow_int32s() {
	out, err := Pow(int32(1), int32(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1
}

func ExamplePow_int64s() {
	out, err := Pow(int64(1), int64(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1
}

func ExamplePow_floats() {
	out, err := Pow(1.11, 2.22)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1.2607152693942678
}
