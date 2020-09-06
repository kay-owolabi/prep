package binarytrees

import (
	"encoding/json"
	"epi/test"
	"strconv"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	testData := test.GetTestData("is_tree_balanced.tsv")[1:]
	type args struct {
		node  *BinaryTreeNode
		array string
	}

	type test struct {
		name string
		args args
		want bool
	}
	var tests []test

	for i, datum := range testData {
		var jsonArray []interface{}
		var res bool

		json.Unmarshal([]byte(datum[0]), &jsonArray)
		json.Unmarshal([]byte(datum[1]), &res)

		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: args{(&BinaryTreeNode{}).BuildBinaryTreeNode(jsonArray), datum[0]},
			want: res,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalanced(tt.args.node); got != tt.want {
				t.Errorf("IsBalanced()\n"+
					"Arguments\n"+
					"\ttree:\t%v\n"+
					"Failure info\n"+
					"\texpected: %v\n"+
					"\tgot: %v\n",
					tt.args.array, tt.want, got)
			}
		})
	}
}
