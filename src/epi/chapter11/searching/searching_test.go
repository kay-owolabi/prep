package searching

import (
	"encoding/json"
	"epi/test"
	"strconv"
	"testing"
)

func TestSearchFirstOfK(t *testing.T) {
	testData := test.GetTestData("search_first_key.tsv")[1:]
	type args struct {
		a []int
		k int
	}

	type test struct {
		name string
		args args
		want int
	}
	var tests []test

	for i, datum := range testData {
		var input []int
		var input1, output int

		json.Unmarshal([]byte(datum[0]), &input)
		json.Unmarshal([]byte(datum[1]), &input1)
		json.Unmarshal([]byte(datum[2]), &output)

		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: args{
				a: input,
				k: input1,
			},
			want: output,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchFirstOfK(tt.args.a, tt.args.k); got != tt.want {
				t.Errorf("SearchFirstOfK() = %v, want %v", got, tt.want)
			}
		})
	}
}
