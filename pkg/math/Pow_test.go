// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerUInt8s(t *testing.T) {
	out, err := Pow(uint8(4), uint8(2))
	assert.NoError(t, err)
	assert.Equal(t, 16, out)
}

func TestPowerZero(t *testing.T) {
	out, err := Pow(int32(8), int32(0))
	assert.Nil(t, err)
	assert.Equal(t, 1, out)
}

func TestPowerNegative(t *testing.T) {
	out, err := Pow(int32(2), int32(-4))
	assert.Nil(t, err)
	assert.Equal(t, 0.0625, out)
}

func TestPowerInts(t *testing.T) {
	out, err := Pow(int32(8), int32(4))
	assert.NoError(t, err)
	assert.Equal(t, 4096, out)
}

func TestPowerInt32s(t *testing.T) {
	out, err := Pow(int32(1), int32(2))
	assert.NoError(t, err)
	assert.Equal(t, 1, out)
}

func TestPowerInt64s(t *testing.T) {
	out, err := Pow(int64(1), int64(2))
	assert.NoError(t, err)
	assert.Equal(t, 1, out)
}

func TestPowerFloats(t *testing.T) {
	out, err := Pow(2, 2.5)
	assert.NoError(t, err)
	assert.Equal(t, 5.65685424949238, out)
}
