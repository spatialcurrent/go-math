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

func TestMeanUInt8s(t *testing.T) {
	in := []uint8{2, 4}
	out, err := Mean(in)
	assert.NoError(t, err)
	assert.Equal(t, float64(3), out)
}

func TestMeanInts(t *testing.T) {
	in := []int{2, 4}
	out, err := Mean(in)
	assert.NoError(t, err)
	assert.Equal(t, 3.0, out)
}

func TestMeanInt32s(t *testing.T) {
	in := []int32{2, 4}
	out, err := Mean(in)
	assert.NoError(t, err)
	assert.Equal(t, 3.0, out)
}

func TestMeanInt64s(t *testing.T) {
	in := []int64{2, 4}
	out, err := Mean(in)
	assert.NoError(t, err)
	assert.Equal(t, 3.0, out)
}

func TestMeanFloats(t *testing.T) {
	in := []float64{1.11, 2.22}
	out, err := Mean(in)
	assert.NoError(t, err)
	assert.Equal(t, 1.665, out)
}

func TestMeanTimes(t *testing.T) {
	in := []time.Duration{
		time.Hour * 2,
		time.Hour * 4,
	}
	out, err := Mean(in)
	assert.NoError(t, err)
	assert.Equal(t, time.Hour*3, out)
}

func TestMeanInterface(t *testing.T) {
	in := []interface{}{
		3,
		1.5,
	}
	out, err := Mean(in)
	assert.NoError(t, err)
	assert.Equal(t, 2.25, out)
}

func TestMeanErrorEmpty(t *testing.T) {
	in := []interface{}{}
	out, err := Mean(in)
	assert.Equal(t, ErrEmptyInput, err)
	assert.Nil(t, out)
}

func TestMeanErrorComparison(t *testing.T) {
	in := []interface{}{
		2,
		4,
		time.Now(),
	}
	out, err := Mean(in)
	assert.IsType(t, &ErrInvalidAddition{}, err)
	assert.Nil(t, out)
}
