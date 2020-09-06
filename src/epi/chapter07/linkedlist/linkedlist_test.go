package linkedlist

import (
	"encoding/json"
	"epi/test"
	"strconv"
	"testing"
)

func GetLinkedList(datum string) *ListNode {
	var parsed []int
	json.Unmarshal([]byte(datum), &parsed)
	LinkedList := FromArray(parsed)
	return LinkedList
}

func TestMergeTwoSortedLists(t *testing.T) {
	testData := test.GetTestData("sorted_lists_merge.tsv")[1:]
	type args struct {
		L1 *ListNode
		L2 *ListNode
	}

	type test struct {
		name string
		args args
		want *ListNode
	}
	var tests []test

	for i, datum := range testData {
		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: args{
				L1: GetLinkedList(datum[0]),
				L2: GetLinkedList(datum[1]),
			},
			want: GetLinkedList(datum[2]),
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoSortedLists(tt.args.L1, tt.args.L2); !got.Equals(tt.want) {
				t.Fatalf("MergeTwoSortedLists()\nFailure info\n\texpected: %s\n\tresult: %s\n",
					tt.want.ToString(), got.ToString())
			}
		})
	}
}
