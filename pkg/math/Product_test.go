// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package math

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProductUInt8s(t *testing.T) {
	in := []uint8{2, 1, 3}
	out, err := Product(in)
	assert.NoError(t, err)
	assert.Equal(t, uint8(6), out)
}

func TestProductInts(t *testing.T) {
	in := []int{2, 1, 3}
	out, err := Product(in)
	assert.NoError(t, err)
	assert.Equal(t, 6, out)
}

func TestProductInt32s(t *testing.T) {
	in := []int32{2, 1, 3}
	out, err := Product(in)
	assert.NoError(t, err)
	assert.Equal(t, int32(6), out)
}

func TestProductInt64s(t *testing.T) {
	in := []int64{2, 1, 3}
	out, err := Product(in)
	assert.NoError(t, err)
	assert.Equal(t, int64(6), out)
}

func TestProductFloats(t *testing.T) {
	in := []float64{1.5, 3.0}
	out, err := Product(in)
	assert.NoError(t, err)
	assert.Equal(t, 4.5, out)
}

func TestProductDuration(t *testing.T) {
	in := []interface{}{
		time.Hour * 1,
		2,
		4,
	}
	out, err := Product(in)
	assert.NoError(t, err)
	assert.Equal(t, time.Hour*8, out)
}

func TestProductInterface(t *testing.T) {
	in := []interface{}{
		1.0,
		2,
		3,
	}
	out, err := Product(in)
	assert.NoError(t, err)
	assert.Equal(t, 6.0, out)
}

func TestProductErrorEmpty(t *testing.T) {
	in := []interface{}{}
	out, err := Product(in)
	assert.Equal(t, ErrEmptyInput, err)
	assert.Nil(t, out)
}

func TestProductErrorComparison(t *testing.T) {
	in := []interface{}{
		2,
		4,
		time.Now(),
	}
	out, err := Product(in)
	assert.IsType(t, &ErrInvalidMultiplication{}, err)
	assert.Nil(t, out)
}
