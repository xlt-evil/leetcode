package main

import (
	"dataproject/Tree"
	"fmt"
)

//递归 -- 自底向上
func maxDepth(root *Tree.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := 0
	rightDepth := 0
	if root.Left != nil {
		leftDepth = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightDepth = maxDepth(root.Right)
	}
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func main() {
	Tree.Result = []int{6, 2, 10, 0, 0, 8, 0, 0, 7, 0, 1}
	tree := Tree.CreateBinaryTree()
	fmt.Println(maxDepth(tree))
	//10 8 2 1 7 6
	//10 8 2 1 6
}
