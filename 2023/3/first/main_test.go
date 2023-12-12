package main

import (
	"testing"
)

type testCaseSolve struct {
	input string
	want  int
}

func TestSolve(T *testing.T) {
	testCases := []testCaseSolve{
		{
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: 4361,
		},
	}

	for _, test := range testCases {
		got := solve(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}
