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
			input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			output: 18,
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
			input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			output: 9,
		},
	}

	for _, test := range testCases {
		assert.DeepEqual(T, test.output, solveSecond(test.input))
	}
}
