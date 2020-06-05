package trees

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) PrintPreOrder() {
	var strBuilder strings.Builder
	t.printPreOrder(&strBuilder)
	fmt.Println("PreOrder: ", strBuilder.String())
}

func (t *TreeNode) PrintInOrder() {
	var strBuilder strings.Builder
	t.printInOrder(&strBuilder)
	fmt.Println("InOrder: ", strBuilder.String())
}

func (t *TreeNode) PrintPostOrder() {
	var strBuilder strings.Builder
	t.printPostOrder(&strBuilder)
	fmt.Println("PostOrder: ", strBuilder.String())
}

func (t *TreeNode) PrintAllOrders() {
	var strBuilder strings.Builder

	t.printPreOrder(&strBuilder)
	fmt.Println("PreOrder: ", strBuilder.String())
	strBuilder.Reset()

	t.printInOrder(&strBuilder)
	fmt.Println("InOrder: ", strBuilder.String())
	strBuilder.Reset()

	t.printPostOrder(&strBuilder)
	fmt.Println("PostOrder: ", strBuilder.String())
	strBuilder.Reset()
}

func (t *TreeNode) printPreOrder(strBuilder *strings.Builder) {
	if t == nil || strBuilder == nil {
		return
	}

	fmt.Fprintf(strBuilder, "%d", t.Val)
	if t.Left != nil {
		fmt.Fprintf(strBuilder, ", ")
		t.Left.printPreOrder(strBuilder)
	}

	if t.Right != nil {
		fmt.Fprintf(strBuilder, ", ")
		t.Right.printPreOrder(strBuilder)
	}
}

func (t *TreeNode) printInOrder(strBuilder *strings.Builder) {
	if t == nil || strBuilder == nil {
		return
	}
	if t.Left != nil {
		t.Left.printInOrder(strBuilder)
		fmt.Fprintf(strBuilder, ", ")
	}
	fmt.Fprintf(strBuilder, "%d", t.Val)
	if t.Right != nil {
		fmt.Fprintf(strBuilder, ", ")
		t.Right.printInOrder(strBuilder)
	}
}

func (t *TreeNode) printPostOrder(strBuilder *strings.Builder) {
	if t == nil || strBuilder == nil {
		return
	}

	if t.Left != nil {
		t.Left.printPostOrder(strBuilder)
		fmt.Fprintf(strBuilder, ", ")
	}
	if t.Right != nil {
		t.Right.printPostOrder(strBuilder)
		fmt.Fprintf(strBuilder, ", ")
	}
	fmt.Fprintf(strBuilder, "%d", t.Val)
}
