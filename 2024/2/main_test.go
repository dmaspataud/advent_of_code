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
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			output: 2,
		},
	}

	for _, test := range testCases {
		assert.DeepEqual(T, test.output, solveFirst(test.input))
	}
}

func TestParseList(T *testing.T) {
	testCases := []struct {
		input  string
		output [][]int
	}{
		{
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			output: [][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
		},
	}

	for _, test := range testCases {
		assert.DeepEqual(T, test.output, parseList(test.input))
	}
}

func TestSolveSecond(T *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			output: 4,
		},
	}

	for _, test := range testCases {
		assert.DeepEqual(T, test.output, solveSecond(test.input))
	}
}
