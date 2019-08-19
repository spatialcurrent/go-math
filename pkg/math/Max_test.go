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

func TestMaxUInt8s(t *testing.T) {
	in := []uint8{2, 1, 3}
	out, err := Max(in)
	assert.NoError(t, err)
	assert.Equal(t, uint8(3), out)
}

func TestMaxInts(t *testing.T) {
	in := []int{2, 1, 3}
	out, err := Max(in)
	assert.NoError(t, err)
	assert.Equal(t, 3, out)
}

func TestMaxInt32s(t *testing.T) {
	in := []int32{2, 1, 3}
	out, err := Max(in)
	assert.NoError(t, err)
	assert.Equal(t, int32(3), out)
}

func TestMaxInt64s(t *testing.T) {
	in := []int64{2, 1, 3}
	out, err := Max(in)
	assert.NoError(t, err)
	assert.Equal(t, int64(3), out)
}

func TestMaxFloats(t *testing.T) {
	in := []float64{2.22, 1.11, 3.33}
	out, err := Max(in)
	assert.NoError(t, err)
	assert.Equal(t, 3.33, out)
}

func TestMaxDuration(t *testing.T) {
	in := []time.Duration{
		time.Hour * 2,
		time.Hour * 3,
		time.Hour * 1,
	}
	out, err := Max(in)
	assert.NoError(t, err)
	assert.Equal(t, time.Hour*3, out)
}

func TestMaxTimes(t *testing.T) {
	now := time.Now()
	in := []time.Time{
		now.Add(time.Hour * -1),
		now,
		now.Add(time.Minute * -2),
	}
	out, err := Max(in)
	assert.NoError(t, err)
	assert.Equal(t, now, out)
}

func TestMaxInterface(t *testing.T) {
	in := []interface{}{
		2,
		3,
		2.22,
		uint8(8),
		int64(2),
		2.0,
	}
	out, err := Max(in)
	assert.NoError(t, err)
	assert.Equal(t, uint8(8), out)
}

func TestMaxErrorEmpty(t *testing.T) {
	in := []interface{}{}
	out, err := Max(in)
	assert.Equal(t, ErrEmptyInput, err)
	assert.Nil(t, out)
}

func TestMaxErrorComparison(t *testing.T) {
	in := []interface{}{
		2,
		4,
		time.Now(),
	}
	out, err := Max(in)
	assert.IsType(t, &ErrInvalidComparison{}, err)
	assert.Nil(t, out)
}
