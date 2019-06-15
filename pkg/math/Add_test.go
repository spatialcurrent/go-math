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

func TestAddUInt8s(t *testing.T) {
	out, err := Add(uint8(1), uint8(2))
	assert.NoError(t, err)
	assert.Equal(t, uint8(3), out)
}

func TestAddInts(t *testing.T) {
	out, err := Add(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 3, out)
}

func TestAddInt32s(t *testing.T) {
	out, err := Add(int32(1), int32(2))
	assert.NoError(t, err)
	assert.Equal(t, int32(3), out)
}

func TestAddInt64s(t *testing.T) {
	out, err := Add(int64(1), int64(2))
	assert.NoError(t, err)
	assert.Equal(t, int64(3), out)
}

func TestAddFloats(t *testing.T) {
	out, err := Add(1.11, 2.22)
	assert.NoError(t, err)
	assert.Equal(t, 3.33, out)
}

func TestAddDurations(t *testing.T) {
	out, err := Add(time.Hour*1, time.Hour*2)
	assert.NoError(t, err)
	assert.Equal(t, time.Hour*3, out)
}

func TestAddTimeDuration(t *testing.T) {
	now := time.Now()
	out, err := Add(now, time.Hour*1)
	assert.NoError(t, err)
	assert.Equal(t, now.Add(time.Hour*1), out)
}
