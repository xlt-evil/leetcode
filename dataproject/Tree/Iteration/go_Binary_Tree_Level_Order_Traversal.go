package main

import (
	"dataproject/Tree"
	"fmt"
)

//树的层序遍历
func levelOrder(root *Tree.TreeNode) [][]int {
	result := [][]int{}
	return DfsLevelOrder(root, 0, result)
}

func DfsLevelOrder(root *Tree.TreeNode, level int, result [][]int) [][]int {
	if root == nil {
		return result
	}
	if level >= len(result) {
		result = append(result, []int{})
	}
	result[level] = append(result[level], root.Val)
	result = DfsLevelOrder(root.Left, level+1, result)
	result = DfsLevelOrder(root.Right, level+1, result)
	return result
}

func BfsLevelOrder(root *Tree.TreeNode) [][]int {
	var result [][]int
	var level int // 几层
	Q := []*Tree.TreeNode{root}
	for len(Q) > 0 {
		length := len(Q)
		level++
		temp := []int{}
		for i := 0; i < length; i++ {
			cur := Q[0]
			Q = Q[1:]
			temp = append(temp, cur.Val)
			if cur.Left != nil {
				Q = append(Q, cur.Left)
			}
			if cur.Right != nil {
				Q = append(Q, cur.Right)
			}
		}
		result = append(result, temp)
	}
	return result
}

func levelOrderBottom(root *Tree.TreeNode) [][]int {
	result := [][]int{}
	result = DfsLevelOrder(root, 0, result)
	length := len(result)
	s := [][]int{}
	for i := length - 1; i >= 0; i-- {
		s = append(s, result[i])
	}
	return s
}

func main() {
	Tree.Result = []int{6, 2, 10, 0, 0, 8, 0, 0, 7, 0, 1}
	tree := Tree.CreateBinaryTree()
	fmt.Println(levelOrderBottom(tree))
	//6 27 1081
}
