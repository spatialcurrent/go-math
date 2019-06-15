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

func TestCompareUInt8s(t *testing.T) {
	out, err := Compare(uint8(1), uint8(2))
	assert.NoError(t, err)
	assert.Equal(t, -1, out)
}

func TestCompareInts(t *testing.T) {
	out, err := Compare(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, -1, out)
}

func TestCompareInt32s(t *testing.T) {
	out, err := Compare(int32(1), int32(2))
	assert.NoError(t, err)
	assert.Equal(t, -1, out)
}

func TestCompareInt64s(t *testing.T) {
	out, err := Compare(int64(1), int64(2))
	assert.NoError(t, err)
	assert.Equal(t, -1, out)
}

func TestCompareFloats(t *testing.T) {
	out, err := Compare(1.11, 2.22)
	assert.NoError(t, err)
	assert.Equal(t, -1, out)
}

func TestCompareDurations(t *testing.T) {
	out, err := Compare(time.Hour*1, time.Hour*2)
	assert.NoError(t, err)
	assert.Equal(t, -1, out)
}

func TestCompareTimes(t *testing.T) {
	now := time.Now()
	out, err := Compare(now, now.Add(time.Hour*1))
	assert.NoError(t, err)
	assert.Equal(t, -1, out)
}
