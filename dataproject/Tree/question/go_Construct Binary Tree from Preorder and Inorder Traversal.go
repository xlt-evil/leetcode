package main

import (
	"dataproject/Tree"
	"fmt"
)

// 知道先序与中序
func buildTree1(preorder []int, inorder []int) *Tree.TreeNode {
	Plength := len(preorder)
	Ilength := len(inorder)
	if Plength == 0 || Ilength == 0 {
		return nil
	}
	root := new(Tree.TreeNode)
	root.Val = preorder[0]
	cl := -1 // 分割线
	for i, _ := range inorder {
		if inorder[i] == root.Val {
			cl = i
			break
		}
	}
	//右子树下标相同
	root.Right = buildTree1(preorder[cl+1:], inorder[cl+1:])
	root.Left = buildTree1(preorder[1:cl+1], inorder[:cl])
	return root
}

func main() {
	tree := buildTree1([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	//tree := buildTree([]int{2,3,1},[]int{3,2,1})
	//tree := buildTree([]int{1,2,3,4},[]int{2,1,4,3})
	//fmt.Println(tree.Left.Right)
	fmt.Println(Tree.DFSInOrder(tree, []int{}))
}
