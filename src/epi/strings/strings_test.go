package strings

import (
	"encoding/json"
	"epi/test"
	"strconv"
	"testing"
)

func TestIntToString(t *testing.T) {
	testData := test.GetTestData("string_integer_interconversion.tsv")[1:]
	type args struct {
		x int
		s string
	}

	type test struct {
		name string
		args args
	}
	var tests []test

	for i, datum := range testData {
		var x int
		var s string
		json.Unmarshal([]byte(datum[0]), &x)
		json.Unmarshal([]byte(datum[1]), &s)

		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: args{
				x: x,
				s: s,
			},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if atoi, err := strconv.Atoi(IntToString(tt.args.x)); err != nil || atoi != tt.args.x {
				t.Errorf("IntToString() = %v, want %v", atoi, tt.args.s)
			}

			if toInt := StringToInt(tt.args.s); toInt != tt.args.x {
				t.Errorf("StringToInt() = %v, want %v", toInt, tt.args.x)
			}
		})
	}
}
