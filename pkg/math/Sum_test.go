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

func TestSumUInt8s(t *testing.T) {
	in := []uint8{2, 1, 3}
	out, err := Sum(in)
	assert.NoError(t, err)
	assert.Equal(t, uint8(6), out)
}

func TestSumInts(t *testing.T) {
	in := []int{2, 1, 3}
	out, err := Sum(in)
	assert.NoError(t, err)
	assert.Equal(t, 6, out)
}

func TestSumInt32s(t *testing.T) {
	in := []int32{2, 1, 3}
	out, err := Sum(in)
	assert.NoError(t, err)
	assert.Equal(t, int32(6), out)
}

func TestSumInt64s(t *testing.T) {
	in := []int64{2, 1, 3}
	out, err := Sum(in)
	assert.NoError(t, err)
	assert.Equal(t, int64(6), out)
}

func TestSumFloats(t *testing.T) {
	in := []float64{4.44, 1.11, 3.33}
	out, err := Sum(in)
	assert.NoError(t, err)
	assert.Equal(t, 8.88, out)
}

func TestSumTimes(t *testing.T) {
	in := []time.Duration{
		time.Hour * 1,
		time.Hour * 2,
	}
	out, err := Sum(in)
	assert.NoError(t, err)
	assert.Equal(t, time.Hour*3, out)
}

func TestSumInterface(t *testing.T) {
	in := []interface{}{
		2,
		3,
		2.22,
		uint8(8),
		int64(2),
		2.0,
	}
	out, err := Sum(in)
	assert.NoError(t, err)
	assert.Equal(t, 19.22, out)
}

func TestSumErrorEmpty(t *testing.T) {
	in := []interface{}{}
	out, err := Sum(in)
	assert.Equal(t, ErrEmptyInput, err)
	assert.Nil(t, out)
}

func TestSumErrorComparison(t *testing.T) {
	in := []interface{}{
		2,
		4,
		time.Now(),
	}
	out, err := Sum(in)
	assert.IsType(t, &ErrInvalidAddition{}, err)
	assert.Nil(t, out)
}
