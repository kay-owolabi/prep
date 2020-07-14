package binarysearchtrees

import (
	"encoding/json"
	"epi/binarytrees"
	"epi/test"
	"testing"
)

func TestIsBinaryTreeBST(t *testing.T) {
	testData := test.GetTestData("is_tree_a_bst.tsv")[1:]
	type args struct {
		tree  *binarytrees.BinaryTreeNode
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
			name: string(i),
			args: args{binarytrees.BuildBinaryTreeNode(jsonArray), datum[0]},
			want: res,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBinaryTreeBST(tt.args.tree); got != tt.want {
				t.Errorf("IsBinaryTreeBST()\n"+
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
