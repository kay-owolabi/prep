package binarytrees

type BinaryTreeNode struct {
	Data        interface{}
	Left, Right *BinaryTreeNode
}

func NewBinaryTree(data interface{}, left, right *BinaryTreeNode) *BinaryTreeNode {
	return &BinaryTreeNode{Data: data, Left: left, Right: right}
}

func BuildBinaryTreeNode(data interface{}) *BinaryTreeNode {
	if data == nil {
		return nil
	}
	var nodes []*BinaryTreeNode
	var root *BinaryTreeNode

	for _, node := range data.([]interface{}) {
		if node == nil {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, NewBinaryTree(node, nil, nil))
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

	return root
}
