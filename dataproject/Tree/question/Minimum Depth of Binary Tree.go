package main

import (
	"dataproject/Tree"
	"fmt"
)

//最小深度
func minDepth(root *Tree.TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 || right == 0 { // 这个地方判断的是单子树的时候
		return 1 + max2(left, right)
	}
	return 1 + min(left, right)
}

//利用Bfs优于Dfs 假设左子树有500个节点，而右子树只有一个节点的情况
func BfsGetMinDepth(root *Tree.TreeNode) int {
	if root == nil {
		return 0
	}
	Q := []*Tree.TreeNode{root}
	length := 0
Done:
	for len(Q) > 0 {
		size := len(Q)
		length++
		for i := 0; i < size; i++ {
			cur := Q[0]
			Q = Q[1:]
			if cur.Right == nil && cur.Left == nil {
				break Done
			}
			if cur.Left != nil {
				Q = append(Q, cur.Left)
			}
			if cur.Right != nil {
				Q = append(Q, cur.Right)
			}
		}
	}
	return length
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	Tree.Result = []int{3, 9, 0, 0, 20, 15, 0, 0, 7}
	tree := Tree.CreateBinaryTree()
	fmt.Println(minDepth(tree))
}
