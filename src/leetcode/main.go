package main

import (
	"fmt"
	"leetcode/fixtures"
	"leetcode/solutions"
)

func main() {
	testFileName := "/Users/koowolab/gitroot/kay-owolabi/prep/src/leetcode/fixtures/invertbinarytree/test.in"
	tree := fixtures.ReadTree(testFileName)
	fmt.Printf("node.Val :%d\n", tree.Val)
	solutions.InvertTree(tree)
}
