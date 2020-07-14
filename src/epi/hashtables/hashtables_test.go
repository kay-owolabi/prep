package hashtables

import (
	"encoding/json"
	"epi/test"
	"testing"
)

func TestIsLetterConstructibleFromMagazine(t *testing.T) {
	testData := test.GetTestData("is_anonymous_letter_constructible.tsv")[1:]
	type args struct {
		letterText   string
		magazineText string
	}
	type test struct {
		name string
		args args
		want bool
	}
	var tests []test
	for i, datum := range testData {
		var input, input1 string
		var output bool

		json.Unmarshal([]byte(datum[0]), &input)
		json.Unmarshal([]byte(datum[1]), &input1)
		json.Unmarshal([]byte(datum[2]), &output)

		tests = append(tests, test{
			name: string(i),
			args: args{
				letterText:   input,
				magazineText: input1,
			},
			want: output,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLetterConstructibleFromMagazine(tt.args.letterText, tt.args.magazineText); got != tt.want {
				t.Errorf("IsLetterConstructibleFromMagazine() = %v, want %v", got, tt.want)
			}
		})
	}
}
