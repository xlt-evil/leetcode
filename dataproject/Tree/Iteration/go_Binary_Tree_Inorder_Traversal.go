package main

import (
	"dataproject/Tree"
	"fmt"
)

// 中序遍历
func inorderTraversal(root *Tree.TreeNode) []int {
	if root == nil {
		return nil
	}
	//result := []int{}
	return BfsInOrder(root)
}

func DFSInOrder(root *Tree.TreeNode, result []int) []int {
	if root != nil {
		// 左
		if root.Left != nil {
			result = DFSInOrder(root.Left, result)
		}
		//根
		result = append(result, root.Val)
		//右
		if root.Right != nil {
			result = DFSInOrder(root.Right, result)
		}
	}
	return result
}

func BfsInOrder(root *Tree.TreeNode) []int {
	result := []int{}
	stack := []*Tree.TreeNode{root}
	rightRoot := []*Tree.TreeNode{} //保存有左节点的根节点
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Left != nil {
			stack = append(stack, cur.Left)
			rightRoot = append(rightRoot, cur)
			continue
		} else {
			result = append(result, cur.Val)
			if cur.Right != nil {
				stack = append(stack, cur.Right)
				continue
			}
			for len(rightRoot) > 0 {
				curs := rightRoot[len(rightRoot)-1]
				rightRoot = rightRoot[:len(rightRoot)-1]
				result = append(result, curs.Val)
				if curs.Right != nil {
					stack = append(stack, curs.Right)
					break
				}
			}
		}
	}
	return result
}

func main() {
	Tree.Result = []int{6, 2, 10, 0, 0, 8, 0, 0, 7, 0, 1}
	tree := Tree.CreateBinaryTree()
	fmt.Println(inorderTraversal(tree))
	//10 2 8 6 7 1
}
