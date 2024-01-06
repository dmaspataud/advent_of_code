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
			input: `.....
.S-7.
.|.|.
.L-J.
.....`,
			want: 4,
		},
		{
			input: `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`,
			want: 8,
		},
		{
			input: `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`,
			want: 4,
		},
	}

	for _, test := range testCases {
		got := solve(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}
