package main

import (
	"dataproject/Tree"
	"fmt"
)

func insertIntoBST(root *Tree.TreeNode, val int) *Tree.TreeNode {
	DfsInsert(root, val)
	return root
}

func BfsInsert(root *Tree.TreeNode, val int) *Tree.TreeNode {
	if root == nil {
		return root
	}
	Q := []*Tree.TreeNode{root}
	prev := root
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == nil {
			break
		}
		prev = cur
		if cur.Val > val {
			Q = append(Q, cur.Left)
			continue
		}
		Q = append(Q, cur.Right)
	}
	next := new(Tree.TreeNode)
	next.Val = val
	if prev.Val > val {
		prev.Left = next
	} else {
		prev.Right = next
	}
	return root
}

func DfsInsert(root *Tree.TreeNode, val int) {
	if root.Val > val {
		if root.Left == nil {
			root.Left = new(Tree.TreeNode)
			root.Left.Val = val
			return
		} else {
			DfsInsert(root.Left, val)
			return
		}
	}
	if root.Right == nil {
		root.Right = new(Tree.TreeNode)
		root.Right.Val = val
		return
	} else {
		DfsInsert(root.Right, val)
		return
	}
}

func main() {
	Tree.Result = []int{4}
	tree := Tree.CreateBinaryTree()
	tree = insertIntoBST(tree, 1)
	fmt.Println(Tree.BfsInOrder(tree))
}
