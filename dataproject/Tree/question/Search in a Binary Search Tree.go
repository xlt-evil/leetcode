package main

import (
	"dataproject/Tree"
	"fmt"
)

func searchBST(root *Tree.TreeNode, val int) *Tree.TreeNode {
	return BfsSearchBST(root, val)
}

func DfsSearchBST(root *Tree.TreeNode, val int) *Tree.TreeNode {
	if root == nil {
		return root
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return DfsSearchBST(root.Left, val)
	} else {
		return DfsSearchBST(root.Right, val)
	}
}

func BfsSearchBST(root *Tree.TreeNode, val int) *Tree.TreeNode {
	if root == nil {
		return root
	}
	Q := []*Tree.TreeNode{root}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == nil {
			return nil
		}
		if cur.Val == val {
			return cur
		}
		if cur.Val > val {
			Q = append(Q, cur.Left)
			continue
		}
		Q = append(Q, cur.Right)
	}
	return nil
}

func main() {
	Tree.Result = []int{4, 2, 1, 0, 0, 3, 0, 0, 7}
	tree := Tree.CreateBinaryTree()
	fmt.Println(searchBST(tree, 2))
	//10 8 2 1 7 6
	//10 8 2 1 6
}
