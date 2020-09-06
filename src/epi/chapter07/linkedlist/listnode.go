package linkedlist

import (
	"fmt"
	"reflect"
	"strings"
)

type ListNode struct {
	data interface{}
	next *ListNode
}

func New(data interface{}, next *ListNode) *ListNode {
	l := new(ListNode)
	(*l).data = data
	(*l).next = next
	return l
}

func FromArray(input interface{}) *ListNode {
	array := reflect.ValueOf(input)
	var head *ListNode
	for i := array.Len() - 1; i >= 0; i-- {
		head = New(array.Index(i).Interface(), head)
	}
	return head
}

func (l *ListNode) ToArray() interface{} {
	var result []interface{}

	for iter := l; iter != nil; iter = iter.next {
		result = append(result, iter.data)
	}

	return result
}

func (l *ListNode) ToString() string {
	result := strings.Builder{}
	visited := map[*ListNode]bool{}
	node := l
	first := true

	for node != nil {
		if first {
			first = false
		} else {
			result.WriteString(" -> ")
		}

		if _, ok := visited[node]; ok {
			if node.next != node {
				result.WriteString(fmt.Sprintf("%v", node.data))
				result.WriteString(" -> ... -> ")
			}
			result.WriteString(fmt.Sprintf("%v", node.data))
			result.WriteString(" -> ...")
			break
		} else {
			result.WriteString(fmt.Sprintf("%v", node.data))
			visited[node] = true
			node = node.next
		}
	}

	return result.String()
}

func (l *ListNode) Equals(other interface{}) bool {
	if l == other.(*ListNode) {
		return true
	}

	if _, ok := other.(*ListNode); other == nil || !ok {
		return false
	}

	return EqualsIterImpl(l, other.(*ListNode))
}

func EqualsIterImpl(a *ListNode, b *ListNode) bool {
	visitedA := make(map[*ListNode]bool)
	visitedB := make(map[*ListNode]bool)

	for a != nil && b != nil {
		if _, ok := visitedA[a]; ok {
			_, ok := visitedB[b]
			return a.data == b.data && ok
		}
		if a.data != b.data {
			return false
		}
		visitedA[a] = true
		visitedB[b] = true

		a = a.next
		b = b.next
	}

	return a == nil && b == nil
}

func (l *ListNode) Size() int {
	var result int
	visited := make(map[*ListNode]bool)
	node := l

	containsNode := false
	for node != nil && !containsNode {
		result++
		visited[node] = true
		node = node.next
		_, containsNode = visited[node]
	}
	return result
}
