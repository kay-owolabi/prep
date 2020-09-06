package arrays

import (
	"encoding/json"
	"epi/test"
	"strconv"
	"testing"
)

func TestDutchFlagPartition(t *testing.T) {
	testData := test.GetTestData("dutch_national_flag.tsv")[1:]

	type args struct {
		pivotIndex int
		colors     []Color
	}
	type test struct {
		name   string
		args   args
		errMsg string
	}
	var tests []test

	for i, datum := range testData {
		var aColors []Color
		var pivotIndex int
		json.Unmarshal([]byte(datum[0]), &aColors)
		json.Unmarshal([]byte(datum[1]), &pivotIndex)
		errMsg := datum[2]

		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: args{
				pivotIndex: pivotIndex,
				colors:     aColors,
			},
			errMsg: errMsg,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var count [3]int
			for _, color := range tt.args.colors {
				count[color]++
			}

			pivot := tt.args.colors[tt.args.pivotIndex]
			DutchFlagPartition(tt.args.pivotIndex, tt.args.colors)

			i := 0
			for i < len(tt.args.colors) && tt.args.colors[i] < pivot {
				count[tt.args.colors[i]]--
				i++
			}

			for i < len(tt.args.colors) && tt.args.colors[i] == pivot {
				count[tt.args.colors[i]]--
				i++
			}

			for i < len(tt.args.colors) && tt.args.colors[i] > pivot {
				count[tt.args.colors[i]]--
				i++
			}

			if i != len(tt.args.colors) {
				t.Logf("Not partitioned after  %dth element\nFailure info\n\tA: %v\n\tpivotIndex: %v\n", i, tt.args.colors, tt.args.pivotIndex)
				t.FailNow()
			} else if count[0] != 0 || count[1] != 0 || count[2] != 0 {
				t.Logf("Some elements are missing from original array\nFailure info\n\tA: %v\n\tpivotIndex: %v\n", tt.args.colors, tt.args.pivotIndex)
				t.FailNow()
			}
		})
	}
}

func TestComputeMaxProfit(t *testing.T) {
	testData := test.GetTestData("buy_and_sell_stock.tsv")[1:]
	type args struct {
		prices []float64
	}
	type test struct {
		name   string
		args   args
		want   float64
		errMsg string
	}
	var tests []test

	for i, datum := range testData {
		var prices []float64
		var want float64
		json.Unmarshal([]byte(datum[0]), &prices)
		json.Unmarshal([]byte(datum[1]), &want)

		tests = append(tests, test{
			name:   strconv.Itoa(i),
			args:   args{prices: prices},
			want:   want,
			errMsg: datum[2],
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeMaxProfit(tt.args.prices); got != tt.want {
				t.Errorf("ComputeMaxProfit() = %v, want %v\nFailure info\n\t: %s\n", got, tt.want, tt.errMsg)
			}
		})
	}
}
