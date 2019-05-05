package main

import (
	"dataproject/Tree"
	"math"
)

func isBalanced(root *Tree.TreeNode) bool {
	return BottomToTop(root) != -1
}

func depth(root *Tree.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

// 先判断最大的左右子树的高度差，在进入下一层节点判断高度差，// 自顶向下===从大到小 从宏观到微观
func TopToBottom(root *Tree.TreeNode) bool {
	if root == nil {
		return true
	}
	left := depth(root.Left) //判断左边的 === 慢慢判断的树会越来越小
	right := depth(root.Right)
	// 但满足左右子树乘积小于等于1时 ,进入到该跟节点的左右子树，判断该跟节点的左右子树的左右子树高度差是不是小于等于1
	return math.Abs(float64(left)-float64(right)) <= 1 && TopToBottom(root.Left) && TopToBottom(root.Right)
}

// 自下向上，从微观到宏观
func BottomToTop(root *Tree.TreeNode) int {
	if root == nil {
		return 0
	}
	//找到最底层的子树，然后去找右子树,
	LeftHight := BottomToTop(root.Left)
	if LeftHight == -1 {
		return -1
	}
	RightHight := BottomToTop(root.Right)
	if RightHight == -1 {
		return -1
	}
	if math.Abs(float64(LeftHight)-float64(RightHight)) > 1 {
		return -1
	}

	return max(RightHight, LeftHight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
