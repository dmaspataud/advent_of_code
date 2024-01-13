package main

import (
	"reflect"
	"testing"
)

type testCaseSolve struct {
	input string
	want  int
}

type testCaseUniverseExpansion struct {
	input [][]string
	want  [][]string
}

type testCaseCalculateDistance struct {
	input [][]int
	want  int
}

func TestSolve(T *testing.T) {
	testCases := []testCaseSolve{
		{
			input: `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`,
			want: 374,
		},
	}

	for _, test := range testCases {
		got := solve(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}

func TestApplyUniverseExpansion(T *testing.T) {
	testCases := []testCaseUniverseExpansion{
		{
			input: [][]string{
				{".", ".", ".", "#", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
				{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
				{".", "#", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
				{"#", ".", ".", ".", "#", ".", ".", ".", ".", "."},
			},
			want: [][]string{
				{".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
				{"#", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", "#", ".", ".", ".", "."},
				{".", "#", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
				{"#", ".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
			},
		},
	}

	for _, test := range testCases {
		got := applyUniverseExpansion(test.input)

		for i := 0; i < len(got); i++ {
			if !reflect.DeepEqual(got[i], test.want[i]) {
				T.Errorf("\nLine %v - Want:\n%v\nGot:\n%v\n", i, test.want[i], got[i])
			}
		}
	}
}

func TestCalculateDistance(T *testing.T) {
	testCases := []testCaseCalculateDistance{
		{
			input: [][]int{{1, 6}, {5, 11}},
			want:  9,
		},
		{
			input: [][]int{{4, 0}, {9, 10}},
			want:  15,
		},
		{
			input: [][]int{{0, 2}, {12, 7}},
			want:  17,
		},
		{
			input: [][]int{{0, 11}, {5, 11}},
			want:  5,
		},
	}

	for _, test := range testCases {
		got := calculateDistance(test.input[0], test.input[1])

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}
