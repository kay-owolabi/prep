package sorting

import (
	"encoding/json"
	"epi/test"
	"reflect"
	"testing"
)

func TestIntersectTwoSortedArrays(t *testing.T) {
	testData := test.GetTestData("intersect_sorted_arrays.tsv")[1:]
	type args struct {
		a, b []int
	}
	type test struct {
		name string
		args args
		want []int
	}
	var tests []test
	for i, datum := range testData {
		var input, input1, output []int
		json.Unmarshal([]byte(datum[0]), &input)
		json.Unmarshal([]byte(datum[1]), &input1)
		json.Unmarshal([]byte(datum[2]), &output)

		tests = append(tests, test{
			name: string(i),
			args: args{
				a: input,
				b: input1,
			},
			want: output,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectTwoSortedArrays(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectTwoSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
