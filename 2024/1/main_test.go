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
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			output: 11,
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
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			output: 31,
		},
	}

	for _, test := range testCases {
		assert.DeepEqual(T, test.output, solveSecond(test.input))
	}
}
