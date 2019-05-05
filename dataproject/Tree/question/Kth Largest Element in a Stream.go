package main

import (
	"fmt"
)

//同时需要用到增加删除查询建议使用二叉树

type KthLargest struct {
	K    int
	Root *TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Count int //记录该节点下有多少个节点
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{K: k, Root: CreateBinaryTree(nums)}
}

func (this *KthLargest) Add(val int) int {
	this.Root = dfsInsert(this.Root, val)
	count := this.K
	temp := this.Root
	for count > 0 {
		pos := 1               // 1 表示当前的点
		if temp.Right != nil { //因为右边总是大于根节点，所以需要先判断右边
			pos += temp.Right.Count
		}
		if count == pos { //当k==pos 就找到了最大K
			return temp.Val
		}
		if count > pos { // 如果最大k数大于pos说明右边的数都在K之前
			count -= pos
			temp = temp.Left //所以最大K值的所在区域在左子树中
		} else { //如果pos 大于 K 就知道 k 存在于右子树中
			temp = temp.Right
		}
	}
	return -1
}

func dfsInsert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = new(TreeNode)
		root.Count = 1
		root.Val = val
		return root
	}
	root.Count++
	if root.Val > val {
		if root.Left == nil {
			root.Left = new(TreeNode)
			root.Left.Val = val
			root.Left.Count = 1
			return root
		} else {
			dfsInsert(root.Left, val)
			return root
		}
	}
	if root.Right == nil {
		root.Right = new(TreeNode)
		root.Right.Val = val
		root.Right.Count = 1
		return root
	} else {
		dfsInsert(root.Right, val)
		return root
	}
}

func CreateBinaryTree(num []int) *TreeNode {
	if len(num) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = num[0]
	root.Count = 1
	for _, v := range num[1:] {
		dfsInsert(root, v)
	}
	return root
}

func main() {
	//myk := Constructor(3,[]int{4,5,8,2})
	//fmt.Println(myk.Add(3))
	//fmt.Println(myk.Add(5))
	//fmt.Println(myk.Add(10))
	//fmt.Println(myk.Add(9))
	//fmt.Println(myk.Add(4))
	myk := Constructor(1, []int{})
	fmt.Println(myk.Add(-3))
	fmt.Println(myk.Add(-2))
	fmt.Println(myk.Add(-4))
	fmt.Println(myk.Add(0))
	fmt.Println(myk.Add(4))
}
