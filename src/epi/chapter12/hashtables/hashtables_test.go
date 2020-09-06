package hashtables

import (
	"encoding/json"
	. "epi/test"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testData := GetTestData("anagrams.tsv")[1:]
	type args struct {
		dictionary []string
	}

	type test struct {
		name string
		args args
		want []StringSlice
	}
	var tests []test
	for i, datum := range testData {
		var input []string
		var output []StringSlice

		json.Unmarshal([]byte(datum[0]), &input)
		json.Unmarshal([]byte(datum[1]), &output)

		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: args{dictionary: input},
			want: output,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < len(tt.want); i++ {
				sort.Strings(tt.want[i])
			}
			sort.Slice(tt.want, func(i, j int) bool {
				return LexicographyComp(tt.want[i], tt.want[j])
			})

			got := FindAnagrams(tt.args.dictionary)
			for i := 0; i < len(got); i++ {
				sort.Strings(got[i])
			}
			sort.Slice(got, func(i, j int) bool {
				return LexicographyComp(StringSlice(got[i]), StringSlice(got[j]))
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nFindAnagrams() = \n got  %v,\n want %v", got, tt.want)
			}
		})
	}
}

func TestIsLetterConstructibleFromMagazine(t *testing.T) {
	testData := GetTestData("is_anonymous_letter_constructible.tsv")[1:]
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
			name: strconv.Itoa(i),
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
