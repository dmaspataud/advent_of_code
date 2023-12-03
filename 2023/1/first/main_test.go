package main

import (
	"testing"
)

type testCase struct {
	input string
	want  int
}

func TestSolve(T *testing.T) {
	testCases := []testCase{
		{
			input: `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`,
			want: 142,
		},
	}

	for _, test := range testCases {
		got := solve(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}
