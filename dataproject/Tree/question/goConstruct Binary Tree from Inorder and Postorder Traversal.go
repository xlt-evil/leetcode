package main

import (
	"dataproject/Tree"
	"fmt"
)

// 知道中序与后续创建一个二叉树

func buildTree(inorder []int, postorder []int) *Tree.TreeNode {
	Ilength := len(inorder)
	Plength := len(postorder)
	// 存在只有单子树的情况
	if Ilength == 0 || Plength == 0 {
		return nil
	}
	cl := -1 // 分割线
	root := new(Tree.TreeNode)
	root.Val = postorder[Plength-1]
	for i, _ := range inorder {
		if inorder[i] == root.Val {
			cl = i
			break
		}
	}
	// 可以知道左子树的范围是一样的
	root.Left = buildTree(inorder[:cl], postorder[:cl])
	root.Right = buildTree(inorder[cl+1:], postorder[cl:Plength-1])
	return root
}

func main() {
	tree := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	//tree := buildTree([]int{2,3,1},[]int{3,2,1})
	//tree := buildTree([]int{1,2,3,4},[]int{2,1,4,3})
	//fmt.Println(tree.Left.Right)
	fmt.Println(Tree.DFSInOrder(tree, []int{}))
}
