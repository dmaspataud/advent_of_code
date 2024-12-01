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
			want: 405,
		},
		{
			input: `..####..##..##..#
...#..##.####.##.
.##.#.##..##..##.
...#..##.####.##.
.##..#..#....#..#
##.##.##########.
#########.##.####`,
			want: 11,
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
				"#.##..#",
				"..##...",
				"##..###",
				"#....#.",
				".#..#.#",
				".#..#.#",
				"#....#.",
				"##..###",
				"..##...",
			},
			want: 5,
		},
		{
			input: []string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			want: 4,
		},
	}

	for _, test := range testCases {
		got := findMirror(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}
