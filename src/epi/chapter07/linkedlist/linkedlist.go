package linkedlist

// 7.1 Merge Two Sorted List
func MergeTwoSortedLists(L1 *ListNode, L2 *ListNode) *ListNode {
	sentinel := new(ListNode)
	it := sentinel

	for L1 != nil && L2 != nil {
		if L1.data.(int) < L2.data.(int) {
			it.next = L1
			L1 = L1.next
		} else {
			it.next = L2
			L2 = L2.next
		}
		it = it.next
	}

	if L1 != nil {
		it.next = L1
	} else {
		it.next = L2
	}

	return sentinel.next
}
