package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
0, 1, 2, 3, 4, 5, 6, 7
0, 1, 1, 2, 3, 5, 8, 13
*/
func TestFibRecursion(t *testing.T) {
	assert.True(t, FibRecursion(5) == 5)
	assert.True(t, FibRecursion(7) == 13)
	assert.True(t, FibRecursion(0) == 0)
	assert.True(t, FibRecursion(1) == 1)
}

func TestFibIteration(t *testing.T) {
	assert.True(t, FibIteration(5) == 5)
	assert.True(t, FibIteration(7) == 13)
	assert.True(t, FibIteration(0) == 0)
	assert.True(t, FibIteration(1) == 1)
}
