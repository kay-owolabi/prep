package greedy

import (
	"encoding/json"
	"epi/test"
	"strconv"
	"testing"
)

func TestChangeMaking(t *testing.T) {
	testData := test.GetTestData("making_change.tsv")[1:]
	type args struct {
		cents int
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
			name: strconv.Itoa(i + 1),
			args: args{cents: input},
			want: output,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeMaking(tt.args.cents); got != tt.want {
				t.Errorf("ChangeMaking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasThreeSum(t *testing.T) {
	testData := test.GetTestData("three_sum.tsv")[1:]
	type args struct {
		A []int
		t int
	}
	type test struct {
		name string
		args args
		want bool
	}
	var tests []test
	for i, datum := range testData {
		var input []int
		var input1 int
		var output bool
		json.Unmarshal([]byte(datum[0]), &input)
		json.Unmarshal([]byte(datum[1]), &input1)
		json.Unmarshal([]byte(datum[2]), &output)

		tests = append(tests, test{
			name: strconv.Itoa(i + 1),
			args: args{
				A: input,
				t: input1,
			},
			want: output,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasThreeSum(tt.args.A, tt.args.t); got != tt.want {
				t.Errorf("HasThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
