package binarysearchtrees

import "math"

func IsBinaryTreeBST(tree *BstNode) bool {
	var isValidBST func(node *BstNode, min, max int) bool

	isValidBST = func(node *BstNode, min, max int) bool {
		if node == nil {
			return true
		} else if min <= node.Data.(int) && node.Data.(int) <= max {
			return isValidBST(node.Left, min, node.Data.(int)) && isValidBST(node.Right, node.Data.(int), max)
		} else {
			return false
		}
	}

	return isValidBST(tree, math.MinInt32, math.MaxInt32)
}

func SearchBST(tree *BstNode, key int) *BstNode {
	if tree == nil || key == tree.Data.(int) {
		return tree
	} else if key < tree.Data.(int) {
		return SearchBST(tree.Left, key)
	} else {
		return SearchBST(tree.Right, key)
	}
}
