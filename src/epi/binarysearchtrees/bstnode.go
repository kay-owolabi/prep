package binarysearchtrees

type BstNode struct {
	Data        interface{}
	Left, Right *BstNode
}

func NewBstNode(data interface{}, left, right *BstNode) *BstNode {
	return &BstNode{Data: data, Left: left, Right: right}
}

func BuildBstNode(data interface{}) *BstNode {
	if data == nil {
		return nil
	}
	var nodes []*BstNode
	var root *BstNode

	for _, node := range data.([]interface{}) {
		if node == nil {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, NewBstNode(node, nil, nil))
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
