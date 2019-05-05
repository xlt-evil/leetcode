package main

import (
	"dataproject/Tree"
	"fmt"
)

// 自底向上，知道子节点答案才能知道答案时使用
func hasPathSum(root *Tree.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if sum == 0 && root.Left == nil && root.Right == nil { //到达根节点
		return true
	}
	flag := hasPathSum(root.Right, sum) || hasPathSum(root.Left, sum)
	return flag
}

func main() {
	Tree.Result = []int{-2, 0, -3}
	tree := Tree.CreateBinaryTree()
	fmt.Println(hasPathSum(tree, -5))
}
