package main

import (
	"testing"
)

type testCaseSolve struct {
	input string
	want  int
}

type testIsValidArrangement struct {
	input  string
	record []int
	want   bool
}

func TestSolve(T *testing.T) {
	testCases := []testCaseSolve{
		{
			input: `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`,
			want: 21,
		},
	}

	for _, test := range testCases {
		got := solve(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}

func TestIsValidArrangement(T *testing.T) {
	testCases := []testIsValidArrangement{
		{
			input:  ".###.##.#...",
			record: []int{3, 2, 1},
			want:   true,
		},
		{
			input:  ".###.##..#..",
			record: []int{3, 2, 1},
			want:   true,
		},
		{
			input:  ".###.##...#.",
			record: []int{3, 2, 1},
			want:   true,
		},
		{
			input:  ".###.##....#",
			record: []int{3, 2, 1},
			want:   true,
		},
		{
			input:  ".###..##.#..",
			record: []int{3, 2, 1},
			want:   true,
		},
		{
			input:  ".###..##..#.",
			record: []int{3, 2, 1},
			want:   true,
		},
		{
			input:  ".###..##...#",
			record: []int{3, 2, 1},
			want:   true,
		},
		{
			input:  ".###...##.#.",
			record: []int{3, 2, 1},
			want:   true,
		},
		{
			input:  ".###...##..#",
			record: []int{3, 2, 1},
			want:   true,
		},
		{
			input:  ".###....##.#",
			record: []int{3, 2, 1},
			want:   true,
		},
		{
			input:  ".###..##.#.#",
			record: []int{3, 2, 1},
			want:   false,
		},
	}

	for _, test := range testCases {
		got := isValidArrangement(test.input, test.record)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}
