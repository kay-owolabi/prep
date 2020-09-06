package binarytrees

func IsBalanced(node *BinaryTreeNode) bool {
	/*if node == nil {
		return true
	}
	return isHeightBalanced(node.Left, node.Right)*/
	return checkBalanced(node).balanced
}

func isHeightBalanced(left *BinaryTreeNode, right *BinaryTreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return right.Left == nil && right.Right == nil
	}
	if right == nil {
		return left.Left == nil && left.Right == nil
	}
	return isHeightBalanced(left.Left, left.Right) && isHeightBalanced(right.Left, right.Right)
}

type BalancedStatusWithHeight struct {
	balanced bool
	height   int
}

func checkBalanced(node *BinaryTreeNode) BalancedStatusWithHeight {
	if node == nil {
		return BalancedStatusWithHeight{
			balanced: true,
			height:   -1,
		}
	}

	leftResult := checkBalanced(node.Left)
	if !leftResult.balanced {
		return leftResult
	}

	rightResult := checkBalanced(node.Right)
	if !rightResult.balanced {
		return rightResult
	}

	return BalancedStatusWithHeight{
		balanced: Abs(leftResult.height-rightResult.height) <= 1,
		height:   Max(leftResult.height, rightResult.height) + 1,
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
