package heaps

import (
	"encoding/json"
	"epi/test"
	"reflect"
	"strconv"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	testData := test.GetTestData("sorted_arrays_merge.tsv")[1:]
	type args struct {
		sortedArrays [][]int
	}
	type test struct {
		name string
		args args
		want []int
	}
	var tests []test
	for i, datum := range testData {
		var input [][]int
		var output []int
		json.Unmarshal([]byte(datum[0]), &input)
		json.Unmarshal([]byte(datum[1]), &output)

		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: args{input},
			want: output,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortedArrays(tt.args.sortedArrays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
