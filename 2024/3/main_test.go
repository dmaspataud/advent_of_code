package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSolveFirst(T *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{
			input:  `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`,
			output: 161,
		},
	}

	for _, test := range testCases {
		assert.DeepEqual(T, test.output, solveFirst(test.input))
	}
}

func TestSolveSecond(T *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{
			input:  `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`,
			output: 48,
		},
	}

	for _, test := range testCases {
		assert.DeepEqual(T, test.output, solveSecond(test.input))
	}
}
