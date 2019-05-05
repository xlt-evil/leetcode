package main

import (
	"dataproject/Tree"
	"fmt"
)

//将排序数组转换为平衡二叉树

func sortedArrayToBST(nums []int) *Tree.TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	root := new(Tree.TreeNode)
	root.Val = nums[length/2]
	root.Left = sortedArrayToBST(nums[:length/2])
	root.Right = sortedArrayToBST(nums[length/2+1:])
	return root
}

func main() {
	tree := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(Tree.BfsInOrder(tree))
}
