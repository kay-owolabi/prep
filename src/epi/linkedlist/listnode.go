package linkedlist

type ListNode struct {
	data interface{}
	next *ListNode
}

func (l *ListNode) New(data interface{}, next *ListNode) {
	(*l).data = data
	(*l).next = next
}

func (l *ListNode) ToArray() []interface{} {
	var result []interface{}

	for iter := l; iter != nil; iter = iter.next {
		result = append(result, iter.data)
	}

	return result
}
