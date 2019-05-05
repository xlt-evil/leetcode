package main

import (
	"dataproject/Tree"
	"fmt"
)

//右旋
//节点都在左子树上
//不平衡点以其左节点为旋转点，向右旋转
func RRotate(root *Tree.TreeNode) { //root 是不平衡点
	L := Tree.TreeNode{}
	L = *root.Left      //获取旋转点
	root.Left = L.Right // 左节点的右节点一定小于节点
	L.Right = root      //向右旋转
}

//左旋
//节点都在右子树上
//不平衡点以其右节点为旋转点，向左旋转
func LRotate(root *Tree.TreeNode) {
	R := root.Right     //获取旋转点
	root.Right = R.Left //右节点的左节点一定大于节点
	R.Left = root       //向左旋转
	*root = *R
}

func main() {
	Tree.Result = []int{3, 1, 0, 0, 6, 4, 0, 5, 0, 0, 7}
	tree := Tree.CreateBinaryTree()
	fmt.Println(Tree.BfsInOrder(tree))
	RRotate(tree)
	fmt.Println(tree.Right)
}
