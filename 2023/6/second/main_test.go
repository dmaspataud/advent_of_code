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
			input: `Time:      7  15   30
Distance:  9  40  200`,
			want: 71503,
		},
	}

	for _, test := range testCases {
		got := solve(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}
