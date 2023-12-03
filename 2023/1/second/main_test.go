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
			input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			want: 281,
		},
	}

	for _, test := range testCases {
		got := solve(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}
