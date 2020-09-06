package binarytrees

type TreeNode interface {
	BuildBinaryTreeNode(data interface{}) *BinaryTreeNode
	NewBinaryTree(data interface{}, left, right *BinaryTreeNode) *BinaryTreeNode
}

type BinaryTreeNode struct {
	Data        interface{}
	Left, Right *BinaryTreeNode
}

func (tree *BinaryTreeNode) NewBinaryTree(data interface{}, left, right *BinaryTreeNode) *BinaryTreeNode {

	*tree = BinaryTreeNode{Data: int(data.(float64)), Left: left, Right: right}
	return tree
}

func (tree *BinaryTreeNode) BuildBinaryTreeNode(data interface{}) *BinaryTreeNode {
	if data == nil {
		return nil
	}
	var nodes []*BinaryTreeNode
	var root *BinaryTreeNode

	for _, node := range data.([]interface{}) {
		if node == nil {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, (&BinaryTreeNode{}).NewBinaryTree(node, nil, nil))
		}
	}

	candidateChildren := nodes
	root, candidateChildren = candidateChildren[0], candidateChildren[1:]

	for _, node := range nodes {
		if node != nil {
			if len(candidateChildren) > 0 {
				node.Left, candidateChildren = candidateChildren[0], candidateChildren[1:]
			}
			if len(candidateChildren) > 0 {
				node.Right, candidateChildren = candidateChildren[0], candidateChildren[1:]
			}
		}
	}

	if root != nil && tree != nil {
		*tree = *root
		return tree
	}
	return nil
}
