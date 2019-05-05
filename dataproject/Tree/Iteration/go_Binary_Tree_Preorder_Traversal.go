package main

import (
	"dataproject/Tree"
	"fmt"
)

//前序遍历 根左右
func preorderTraversal(root *Tree.TreeNode) []int {
	if root == nil {
		return nil
	}
	result := BfsMyPre(root)
	return result
}

func DfsMyPre(root *Tree.TreeNode, result []int) []int {
	//根
	if root != nil {
		result = append(result, root.Val)
		//左
		if root.Left != nil {
			result = DfsMyPre(root.Left, result)
		}
		// 右
		if root.Right != nil {
			result = DfsMyPre(root.Right, result)
		}
	}
	return result
}

func BfsMyPre(root *Tree.TreeNode) []int {
	result := []int{}
	stack := []*Tree.TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return result
}

func main() {
	Tree.Result = []int{6, 2, 10, 0, 0, 8, 0, 0, 7, 0, 1}
	tree := Tree.CreateBinaryTree()
	fmt.Println(preorderTraversal(tree))
}
