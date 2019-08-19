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

func TestMinUInt8s(t *testing.T) {
	in := []uint8{2, 1, 3}
	out, err := Min(in)
	assert.NoError(t, err)
	assert.Equal(t, uint8(1), out)
}

func TestMinInts(t *testing.T) {
	in := []int{2, 1, 3}
	out, err := Min(in)
	assert.NoError(t, err)
	assert.Equal(t, 1, out)
}

func TestMinInt32s(t *testing.T) {
	in := []int32{2, 1, 3}
	out, err := Min(in)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), out)
}

func TestMinInt64s(t *testing.T) {
	in := []int64{2, 1, 3}
	out, err := Min(in)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), out)
}

func TestMinFloats(t *testing.T) {
	in := []float64{2.22, 1.11, 3.33}
	out, err := Min(in)
	assert.NoError(t, err)
	assert.Equal(t, 1.11, out)
}

func TestMinDuration(t *testing.T) {
	in := []time.Duration{
		time.Hour * 2,
		time.Hour * 3,
		time.Hour * 1,
	}
	out, err := Min(in)
	assert.NoError(t, err)
	assert.Equal(t, time.Hour*1, out)
}

func TestMinTimes(t *testing.T) {
	now := time.Now()
	in := []time.Time{
		now.Add(time.Hour * 1),
		now,
		now.Add(time.Minute * 2),
	}
	out, err := Min(in)
	assert.NoError(t, err)
	assert.Equal(t, now, out)
}

func TestMinInterface(t *testing.T) {
	in := []interface{}{
		2,
		3,
		2.22,
		uint8(8),
		int64(2),
		2.0,
	}
	out, err := Min(in)
	assert.NoError(t, err)
	assert.Equal(t, 2, out)
}

func TestMinErrorEmpty(t *testing.T) {
	in := []interface{}{}
	out, err := Min(in)
	assert.Equal(t, ErrEmptyInput, err)
	assert.Nil(t, out)
}

func TestMinErrorComparison(t *testing.T) {
	in := []interface{}{
		2,
		4,
		time.Now(),
	}
	out, err := Min(in)
	assert.IsType(t, &ErrInvalidComparison{}, err)
	assert.Nil(t, out)
}
