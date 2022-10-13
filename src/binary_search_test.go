package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {

	assert.True(t, math.Abs(math.Sqrt(5)-Sqrt(5)) < 0.001)
}
