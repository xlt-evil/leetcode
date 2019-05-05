package main

import (
	"dataproject/Tree"
	"fmt"
)

func postorderTraversal(root *Tree.TreeNode) []int {
	if root != nil {
		return nil
	}
	return BfsPostOrder(root)
}

func DfsPostOrder(root *Tree.TreeNode, result []int) []int {
	if root != nil {
		//左
		if root.Left != nil {
			result = DfsPostOrder(root.Left, result)
		}
		if root.Right != nil {
			result = DfsPostOrder(root.Right, result)
		}
		result = append(result, root.Val)
	}
	return result
}

func BfsPostOrder(root *Tree.TreeNode) []int {
	result := []int{}
	stack := []*Tree.TreeNode{root}
	LeftRoot := []*Tree.TreeNode{} //保存有左节点的根节点
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Left != nil {
			stack = append(stack, cur.Left)
			LeftRoot = append(LeftRoot, cur)
			continue
		} else {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
				//该节点表示当前节点左右节点都遍历完后的临时存放点
				temp := new(Tree.TreeNode)
				temp.Val = cur.Val
				LeftRoot = append(LeftRoot, temp)
				continue
			}
			result = append(result, cur.Val)
			for len(LeftRoot) > 0 {
				curs := LeftRoot[len(LeftRoot)-1]
				LeftRoot = LeftRoot[:len(LeftRoot)-1]
				if curs.Right != nil {
					stack = append(stack, curs.Right)
					//该节点表示当前节点左右节点都遍历完后的临时存放点
					temp := new(Tree.TreeNode)
					temp.Val = curs.Val
					LeftRoot = append(LeftRoot, temp)
					break
				}
				result = append(result, curs.Val)
			}
		}
	}
	return result
}

func main() {
	Tree.Result = []int{6, 2, 10, 0, 0, 8, 0, 0, 7, 0, 1}
	tree := Tree.CreateBinaryTree()
	fmt.Println(postorderTraversal(tree))
	//10 8 2 1 7 6
	//10 8 2 1 6
}
