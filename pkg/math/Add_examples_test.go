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

func ExampleAdd_uInt8s() {
	out, err := Add(uint8(1), uint8(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 3
}

func ExampleAdd_ints() {
	out, err := Add(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 3
}

func ExampleAdd_int32s() {
	out, err := Add(int32(1), int32(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 3
}

func ExampleAdd_int64s() {
	out, err := Add(int64(1), int64(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 3
}

func ExampleAdd_floats() {
	out, err := Add(1.11, 2.22)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 3.33
}

func ExampleAdd_durations() {
	out, err := Add(time.Hour*1, time.Hour*2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 3h0m0s
}
