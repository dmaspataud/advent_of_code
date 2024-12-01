package main

import (
	"testing"
)

type testCaseSolve struct {
	input string
	want  int
}

type testFindMirror struct {
	input []string
	want  int
}

type testFindSmudge struct {
	input1 string
	input2 string
	want   bool
}

func TestSolve(T *testing.T) {
	testCases := []testCaseSolve{
		{
			input: `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`,
			want: 400,
		},
	}

	for _, test := range testCases {
		got := solve(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}

func TestFindMirror(T *testing.T) {
	testCases := []testFindMirror{
		{
			input: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
			want: 3,
		},
		{
			input: []string{
				"#...##..#",
				"#....#..#", // edge case, a one line mirror with smudge would not work with current algo -> we need to inspect line 0 to 1 and last to last - 1
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			want: 1,
		},
	}

	for _, test := range testCases {
		got := findMirror(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}

func TestFindSmudge(T *testing.T) {
	testCases := []testFindSmudge{
		{
			input1: "",
			input2: "",
			want:   false,
		},
		{
			input1: "#.##..##.",
			input2: "..##..##.",
			want:   true,
		},
		{
			input1: "#.#.##.#.",
			input2: "#####.###",
			want:   false,
		},
		{
			input1: "#####.##.",
			input2: "",
			want:   false,
		},
	}

	for _, test := range testCases {
		got := findSmudge(test.input1, test.input2)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}
