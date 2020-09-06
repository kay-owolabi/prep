package main

import (
	"fmt"
)

func main() {
	//fmt.Printf("Forest: %v\nDeleted: %v\n", val, DeleteNode(val, 3))

	//parallelcomputing.OddEvenMonitor()
	input := []string{"alpha", "beta", "gamma"}
	rules := [][]string{{"alpha", "zeta"}, {"beta", "omega"}, {"gamma", "alpha"}}

	SubStrings(input, rules)
	fmt.Println(input)
	//arrays.SpiralOutwards(20)
}

var val = Forest{
	TreeNode{0, ""},
	TreeNode{0, ""},
	TreeNode{0, ""},
	TreeNode{1, ""},
	TreeNode{1, ""},
	TreeNode{4, ""},
	TreeNode{4, ""},
	TreeNode{2, ""},
	TreeNode{2, ""},
	TreeNode{9, ""},
}

type TreeNode struct {
	parent int
	data   string
}

type Forest []TreeNode

func DeleteNode(forest Forest, nodeToDelete int) Forest {
	result := Forest{}
	newParent := map[int]int{}
	for i, node := range forest {
		if i == nodeToDelete || forest[node.parent].parent == -1 {
			forest[i].parent = -1
		} else {
			newParent[i] = len(result)
			node.parent = newParent[node.parent]
			result = append(result, node)
		}
	}
	return result
}

/*

Input: ["alpha", "beta", "gamma", …]
Substitution Rules:  [("alpha", "zeta"), ("beta", "omega"), ("gamma", "alpha")].

*/
