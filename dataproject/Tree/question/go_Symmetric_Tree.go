package main

import (
	"dataproject/Tree"
	"fmt"
)

//判断二叉树是否是对称二叉树
//自顶向下 --答案是肯定的例如结构必须是肯定的
func isSymmetric(root *Tree.TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Right == nil && root.Left == nil {
		return true
	}
	if root.Right == nil || root.Left == nil {
		return false
	}

	return isSymmetry(root.Left, root.Right)
}

func isSymmetry(n1, n2 *Tree.TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	if n1.Val != n2.Val {
		return false
	}
	return isSymmetry(n1.Left, n2.Right) && isSymmetry(n1.Right, n2.Left)
}

func main() {
	Tree.Result = []int{1, 2, 3, 0, 0, 4, 0, 0, 2, 4, 0, 0, 3}
	//Tree.Result = []int{1,2,0,3,0,0,2,0,3}
	//Tree.Result = []int{1}
	tree := Tree.CreateBinaryTree()
	fmt.Println(isSymmetric(tree))
}
