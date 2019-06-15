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

func ExampleCompare_uInt8s() {
	out, err := Compare(uint8(1), uint8(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: -1
}

func ExampleCompare_ints() {
	out, err := Compare(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: -1
}

func ExampleCompare_int32s() {
	out, err := Compare(int32(1), int32(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: -1
}

func ExampleCompare_int64s() {
	out, err := Compare(int64(1), int64(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: -1
}

func ExampleCompare_floats() {
	out, err := Compare(1.11, 2.22)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: -1
}

func ExampleCompare_durations() {
	out, err := Compare(time.Hour*1, time.Hour*2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: -1
}

func ExampleCompare_times() {
	now := time.Now()
	out, err := Compare(now, now.Add(time.Hour*1))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
