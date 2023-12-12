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
			input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
			want: 6440,
		},
	}

	for _, test := range testCases {
		got := solve(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}
