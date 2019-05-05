package main

import (
	"dataproject/Tree"
)

//二叉搜索树
//节点必须大于等于左子树，小于等于右子树

//验证是否是二叉树，最有效的方法是中序遍历 ,如果是递增的则说明是

func isValidBST(root *Tree.TreeNode) bool {
	if root == nil {
		return true
	}
	result := Tree.BfsInOrder(root)
	length := len(result)
	for i := 0; i < length-1; i++ {
		if i+1 < length {
			if result[i] > result[i+1] {
				return false
			}
		}
	}
	return true
}
