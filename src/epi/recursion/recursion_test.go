package recursion

import (
	"encoding/json"
	"epi/test"
	"reflect"
	"testing"
)

func TestComputeTowerHanoi(t *testing.T) {
	testData := test.GetTestData("hanoi.tsv")[1:]
	type args struct {
		numRings int
	}
	type test struct {
		name string
		args args
		want [][]int
	}
	var tests []test
	for i, datum := range testData {
		var input int
		json.Unmarshal([]byte(datum[0]), &input)

		tests = append(tests, test{
			name: string(i),
			args: args{numRings: input},
			// TODO: Finish this test function
			want: nil,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeTowerHanoi(tt.args.numRings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeTowerHanoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
