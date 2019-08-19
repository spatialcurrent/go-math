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

func TestSubtractUInt8s(t *testing.T) {
	out, err := Subtract(uint8(4), uint8(2))
	assert.NoError(t, err)
	assert.Equal(t, uint8(2), out)
}

func TestSubtractInts(t *testing.T) {
	out, err := Subtract(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, -1, out)
}

func TestSubtractInt32s(t *testing.T) {
	out, err := Subtract(int32(1), int32(2))
	assert.NoError(t, err)
	assert.Equal(t, int32(-1), out)
}

func TestSubtractInt64s(t *testing.T) {
	out, err := Subtract(int64(1), int64(2))
	assert.NoError(t, err)
	assert.Equal(t, int64(-1), out)
}

func TestSubtractFloats(t *testing.T) {
	out, err := Subtract(1.11, 2.22)
	assert.NoError(t, err)
	assert.Equal(t, -1.11, out)
}

func TestSubtractDurations(t *testing.T) {
	out, err := Subtract(time.Hour*3, time.Hour*1)
	assert.NoError(t, err)
	assert.Equal(t, time.Hour*2, out)
}

func TestSubtractTimeDuration(t *testing.T) {
	now := time.Now()
	out, err := Subtract(now, time.Hour*1)
	assert.NoError(t, err)
	assert.Equal(t, now.Add(time.Duration(int64(-1)*int64(time.Hour*1))), out)
}
