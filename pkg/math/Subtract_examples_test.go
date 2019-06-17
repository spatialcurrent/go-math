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

func ExampleSubtract_uInt8s() {
	out, err := Subtract(uint8(4), uint8(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 2
}

func ExampleSubtract_ints() {
	out, err := Subtract(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: -1
}

func ExampleSubtract_int32s() {
	out, err := Subtract(int32(1), int32(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: -1
}

func ExampleSubtract_int64s() {
	out, err := Subtract(int64(1), int64(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: -1
}

func ExampleSubtract_floats() {
	out, err := Subtract(1.11, 2.22)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: -1.11
}

func ExampleSubtract_durations() {
	out, err := Subtract(time.Hour*3, time.Hour*2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 1h0m0s
}

func ExampleSubtract_timeDuration() {
	out, err := Subtract(time.Now(), time.Hour*2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
