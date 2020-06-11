package primitivetypes

import (
	"epi/test"
	"strconv"
	"testing"
)

func TestParity(t *testing.T) {
	testData := test.GetTestData("parity.tsv")[1:]

	type args struct {
		x uint64
	}
	var tests []struct {
		name       string
		args       args
		want       uint16
		errMessage string
	}
	for i, datum := range testData {
		name := strconv.Itoa(i + 1)
		x, _ := strconv.ParseUint(datum[0], 10, 64)
		want, _ := strconv.ParseUint(datum[1], 10, 16)
		errMessage := datum[2]

		tests = append(tests, struct {
			name       string
			args       args
			want       uint16
			errMessage string
		}{name: name, args: args{x: x}, want: uint16(want), errMessage: errMessage})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parity(tt.args.x); got != tt.want {
				t.Logf("Parity() = %v, want %v\nFailure info\n\t: %s\n", got, tt.want, tt.errMessage)
				t.FailNow()
			}
		})
	}
}
