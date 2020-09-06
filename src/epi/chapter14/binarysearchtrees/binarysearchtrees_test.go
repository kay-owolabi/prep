package binarysearchtrees

import (
	"encoding/json"
	"epi/binarytrees"
	"epi/test"
	"reflect"
	"strconv"
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
			name: strconv.Itoa(i),
			args: args{(&BstNode{}).BuildBinaryTreeNode(jsonArray), datum[0]},
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

func TestSearchBST(t *testing.T) {
	testData := test.GetTestData("search_in_bst.tsv")[1:]
	type args struct {
		tree *BstNode
		key  int
	}

	type test struct {
		name string
		args args
		want int
	}
	var tests []test

	for i, datum := range testData {
		var jsonArray []interface{}
		var input int
		var output int

		json.Unmarshal([]byte(datum[0]), &jsonArray)
		json.Unmarshal([]byte(datum[1]), &input)
		json.Unmarshal([]byte(datum[2]), &output)

		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: args{
				tree: (&BstNode{}).BuildBinaryTreeNode(jsonArray),
				key:  input,
			},
			want: output,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := SearchBST(tt.args.tree, tt.args.key)
			var got int
			if ret == nil {
				got = -1
			} else {
				got = ret.Data.(int)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
