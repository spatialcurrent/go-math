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
)

import (
	"github.com/stretchr/testify/assert"
)

func TestDivideUInt8s(t *testing.T) {
	out, err := Divide(uint8(4), uint8(2))
	assert.NoError(t, err)
	assert.Equal(t, uint8(2), out)
}

func TestDivideInts(t *testing.T) {
	out, err := Divide(int32(8), int32(4))
	assert.NoError(t, err)
	assert.Equal(t, int32(2), out)
}

func TestDivideInt32s(t *testing.T) {
	out, err := Divide(int32(1), int32(2))
	assert.NoError(t, err)
	assert.Equal(t, 0.5, out)
}

func TestDivideInt64s(t *testing.T) {
	out, err := Divide(int64(1), int64(2))
	assert.NoError(t, err)
	assert.Equal(t, 0.5, out)
}

func TestDivideFloats(t *testing.T) {
	out, err := Divide(1.11, 2.22)
	assert.NoError(t, err)
	assert.Equal(t, 0.5, out)
}

func TestDivideDurations(t *testing.T) {
	out, err := Divide(time.Hour*3, time.Hour*2)
	assert.NoError(t, err)
	assert.Equal(t, 1.5, out)
}

func TestDivideDurationInt(t *testing.T) {
	out, err := Divide(time.Hour*4, 2)
	assert.NoError(t, err)
	assert.Equal(t, time.Hour*2, out)
}
