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

func ExampleProduct_uint8s() {
	out, err := Product([]uint8{1, 2, 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 6
}

func ExampleProduct_ints() {
	out, err := Product([]int{10, 20})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 200
}

func ExampleProduct_floats() {
	out, err := Product([]float64{1.5, 3.0})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 4.5
}

func ExampleProduct_duration() {
	out, err := Product([]interface{}{
		time.Hour * 1,
		2,
		4,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 8h0m0s
}

func ExampleProduct_interface() {
	out, err := Product([]interface{}{10, 2.5})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 25
}
