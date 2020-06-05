package solutions

import "leetcode/list"

func deleteNode(node *list.Node) {
	*node = *(node.Next)
}
