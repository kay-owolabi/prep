package recursion

import (
	"container/list"
	"encoding/json"
	. "epi/test"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestComputeTowerHanoi(t *testing.T) {
	testData := GetTestData("hanoi.tsv")[1:]
	type args struct {
		numRings int
	}
	type test struct {
		name string
		args args
	}
	var tests []test
	for i, datum := range testData {
		var input int
		json.Unmarshal([]byte(datum[0]), &input)

		tests = append(tests, test{
			name: strconv.Itoa(i + 1),
			args: args{numRings: input},
		})
	}
	const NumPegs = 3
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pegs []*list.List
			for i := 0; i < NumPegs; i++ {
				pegs = append(pegs, list.New())
			}
			for i := tt.args.numRings; i >= 1; i-- {
				pegs[0].PushFront(i)
			}
			result := ComputeTowerHanoi(tt.args.numRings)
			fmt.Printf("result : %v len: %v\n", result, len(result))
			for _, operation := range result {
				fromPeg := operation[0]
				toPeg := operation[1]
				if pegs[toPeg].Len() != 0 && pegs[fromPeg].Front().Value.(int) >= pegs[toPeg].Front().Value.(int) {
					t.Errorf("Illegal move from %v to %v", pegs[fromPeg].Front().Value, pegs[toPeg].Front().Value)
					return
				}
				pegs[toPeg].PushFront(pegs[fromPeg].Remove(pegs[fromPeg].Front()))
			}

			var expectedPegs1 []*list.List
			for i := 0; i < NumPegs; i++ {
				expectedPegs1 = append(expectedPegs1, list.New())
			}
			for i := tt.args.numRings; i >= 1; i-- {
				expectedPegs1[1].PushFront(i)
			}

			var expectedPegs2 []*list.List
			for i := 0; i < NumPegs; i++ {
				expectedPegs2 = append(expectedPegs2, list.New())
			}
			for i := tt.args.numRings; i >= 1; i-- {
				expectedPegs2[2].PushFront(i)
			}

			if !reflect.DeepEqual(pegs, expectedPegs1) && !reflect.DeepEqual(pegs, expectedPegs2) {
				t.Errorf("ComputeTowerHanoi() = Pegs doesn't place in the right configuration")
			}
		})
	}
}

func TestGCD(t *testing.T) {
	testData := GetTestData("gcd.tsv")[1:]
	type args struct {
		x int
		y int
	}
	type test struct {
		name string
		args args
		want int
	}
	var tests []test

	for i, datum := range testData {
		var input, input1, output int
		json.Unmarshal([]byte(datum[0]), &input)
		json.Unmarshal([]byte(datum[1]), &input1)
		json.Unmarshal([]byte(datum[2]), &output)

		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: args{
				x: input,
				y: input1,
			},
			want: output,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNQueens(t *testing.T) {
	testData := GetTestData("n_queens.tsv")[1:]
	type args struct {
		n int
	}

	type test struct {
		name string
		args args
		want []IntSlice
	}

	var tests []test

	for i, datum := range testData {
		var input int
		var output []IntSlice
		json.Unmarshal([]byte(datum[0]), &input)
		json.Unmarshal([]byte(datum[1]), &output)

		tests = append(tests, test{
			name: strconv.Itoa(i + 1),
			args: args{input},
			want: output,
		})
	}

	comp := func(expected, result []IntSlice) bool {
		if len(result) == 0 {
			return len(expected) == 0
		}

		sort.Slice(expected, func(i, j int) bool {
			return LexicographyComp(expected[i], expected[j])
		})
		sort.Slice(result, func(i, j int) bool {
			return LexicographyComp(result[i], result[j])
		})
		return reflect.DeepEqual(result, expected)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NQueens(tt.args.n); !comp(got, tt.want) {
				t.Errorf("NQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
