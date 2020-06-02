package solutions

import "leetcode/trees"

func InvertTree(root *trees.TreeNode) *trees.TreeNode {
	if root == nil {
		return root
	}

	if root.Left != nil || root.Right != nil {
		root.Left, root.Right = root.Right, root.Left
		InvertTree(root.Left)
		InvertTree(root.Right)
	}
	return root
}
