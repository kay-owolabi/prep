package dynamicprogramming

import (
	"encoding/json"
	"epi/test"
	"strconv"
	"testing"
)

func TestFibonacci(t *testing.T) {
	testData := test.GetTestData("fibonacci.tsv")[1:]
	type args struct {
		n int
	}
	type test struct {
		name string
		args args
		want int
	}
	var tests []test
	for i, datum := range testData {
		var input, output int
		json.Unmarshal([]byte(datum[0]), &input)
		json.Unmarshal([]byte(datum[1]), &output)

		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: args{
				n: input,
			},
			want: output,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaximumSubarray(t *testing.T) {
	testData := test.GetTestData("max_sum_subarray.tsv")[1:]
	type args struct {
		A []int
	}
	type test struct {
		name string
		args args
		want int
	}

	var tests []test

	for i, datum := range testData {
		var input []int
		var output int
		json.Unmarshal([]byte(datum[0]), &input)
		json.Unmarshal([]byte(datum[1]), &output)

		tests = append(tests, test{
			name: strconv.Itoa(i + 1),
			args: args{A: input},
			want: output,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaximumSubarray(tt.args.A); got != tt.want {
				t.Errorf("FindMaximumSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumCombinationsForFinalScore(t *testing.T) {
	testData := test.GetTestData("number_of_score_combinations.tsv")[1:]
	type args struct {
		finalScore           int
		individualPlayScores []int
	}
	type test struct {
		name string
		args args
		want int
	}
	var tests []test

	for i, datum := range testData {
		var input0 int
		var input1 []int
		var output int
		json.Unmarshal([]byte(datum[0]), &input0)
		json.Unmarshal([]byte(datum[1]), &input1)
		json.Unmarshal([]byte(datum[2]), &output)

		tests = append(tests, test{
			name: strconv.Itoa(i + 1),
			args: args{
				finalScore:           input0,
				individualPlayScores: input1,
			},
			want: output,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumCombinationsForFinalScore(tt.args.finalScore, tt.args.individualPlayScores); got != tt.want {
				t.Errorf("NumCombinationsForFinalScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
