package main

import (
	"reflect"
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

type testUnfold struct {
	input      string
	record     []int
	wantInput  string
	wantRecord []int
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
			want: 525152,
		},
	}

	for _, test := range testCases {
		got := solve(test.input)

		if test.want != got {
			T.Errorf("Want %v, got %v", test.want, got)
		}
	}
}

func TestUnfold(T *testing.T) {
	testCases := []testUnfold{
		{
			input:      ".#",
			record:     []int{1},
			wantInput:  ".#?.#?.#?.#?.#",
			wantRecord: []int{1, 1, 1, 1, 1},
		},
		{
			input:      "???.###",
			record:     []int{1, 1, 3},
			wantInput:  "???.###????.###????.###????.###????.###",
			wantRecord: []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3},
		},
	}

	for _, test := range testCases {
		got := unfold(springLine{springs: test.input, record: test.record})

		if test.wantInput != got.springs {
			T.Errorf("Want %v, got %v", test.wantInput, got.springs)
		}
		if !reflect.DeepEqual(test.wantRecord, got.record) {
			T.Errorf("Want %v, got %v", test.wantRecord, got.record)
		}
	}
}

type testSum struct {
	numbers []int
	sum     int
}

func TestSum(T *testing.T) {
	testCases := []testSum{
		{
			numbers: []int{2, 2},
			sum:     4,
		},
		{
			numbers: []int{2, 2, 2},
			sum:     6,
		},
		{
			numbers: []int{-50, 2},
			sum:     -48,
		},
		{
			numbers: []int{3},
			sum:     3,
		},
		{
			numbers: []int{},
			sum:     0,
		},
	}

	for _, test := range testCases {
		got := sum(test.numbers)
		if got != test.sum {
			T.Errorf("Want %v, got %v", test.sum, got)
		}
	}
}
