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

func TestMultiplyUInt8s(t *testing.T) {
	out, err := Multiply(uint8(1), uint8(2))
	assert.NoError(t, err)
	assert.Equal(t, uint8(2), out)
}

func TestMultiplyInts(t *testing.T) {
	out, err := Multiply(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, out)
}

func TestMultiplyInt32s(t *testing.T) {
	out, err := Multiply(int32(1), int32(2))
	assert.NoError(t, err)
	assert.Equal(t, int32(2), out)
}

func TestMultiplyInt64s(t *testing.T) {
	out, err := Multiply(int64(1), int64(2))
	assert.NoError(t, err)
	assert.Equal(t, int64(2), out)
}

func TestMultiplyFloats(t *testing.T) {
	out, err := Multiply(4, 2)
	assert.NoError(t, err)
	assert.Equal(t, 8, out)
}

func TestMultiplyDurationInt(t *testing.T) {
	out, err := Multiply(2, time.Hour*1)
	assert.NoError(t, err)
	assert.Equal(t, time.Hour*2, out)
}

func TestMultiplyDurationFloat(t *testing.T) {
	expected, err := time.ParseDuration("1h30m")
	assert.NoError(t, err)
	out, err := Multiply(1.5, time.Hour*1)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
