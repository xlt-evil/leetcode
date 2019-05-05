package main

import (
	"dataproject/Tree"
	"fmt"
)

//同时执行搜索、插入、删除
// 根节点的左子树一定都小于跟
// 根节点的右子树一定都大于跟
func deleteNode(root *Tree.TreeNode, key int) *Tree.TreeNode {
	if root == nil {
		return root
	}
	//先找到对应删除节点
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else { // 找到了对应的删除点
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		//存在左右子树的情况，选择去寻找右子树最小的点或左子树最大的点
		//这里选择去寻找右子树最小的点
		minnode := findMinAndDel(root.Right)
		root.Val = minnode.Val
		root.Right = deleteNode(root.Right, root.Val)
		//替换值
	}
	return root
}

//最左的点是最小的点
func findMinAndDel(root *Tree.TreeNode) *Tree.TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func main() {
	Tree.Result = []int{5, 3, 2, 0, 0, 4, 0, 0, 6, 0, 7}
	tree := Tree.CreateBinaryTree()
	tree = deleteNode(tree, 3)
	fmt.Println(Tree.BfsInOrder(tree))
}
