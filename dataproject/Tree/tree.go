package Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var Result []int

//切片也是值创递，只不过是因为共有了同一个地址块
func CreateBinaryTree() *TreeNode {
	if len(Result) == 0 {
		return nil
	}
	val := Result[0]
	Result = Result[1:]
	var root *TreeNode
	if val != 0 {
		root = new(TreeNode)
		root.Val = val
		// 先创建左子树
		root.Left = CreateBinaryTree()
		// 创建右子树
		root.Right = CreateBinaryTree()
	}
	return root
}

func DFSInOrder(root *TreeNode, result []int) []int {
	if root != nil {
		// 左
		if root.Left != nil {
			result = DFSInOrder(root.Left, result)
		}
		//根
		result = append(result, root.Val)
		//右
		if root.Right != nil {
			result = DFSInOrder(root.Right, result)
		}
	}
	return result
}

func DfsMyPre(root *TreeNode, result []int) []int {
	//根
	if root != nil {
		result = append(result, root.Val)
		//左
		if root.Left != nil {
			result = DfsMyPre(root.Left, result)
		}
		// 右
		if root.Right != nil {
			result = DfsMyPre(root.Right, result)
		}
	}
	return result
}

func BfsMyPre(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return result
}

func BfsInOrder(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{root}
	rightRoot := []*TreeNode{} //保存有左节点的根节点
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Left != nil {
			stack = append(stack, cur.Left)
			rightRoot = append(rightRoot, cur)
			continue
		} else {
			result = append(result, cur.Val)
			if cur.Right != nil {
				stack = append(stack, cur.Right)
				continue
			}
			for len(rightRoot) > 0 {
				curs := rightRoot[len(rightRoot)-1]
				rightRoot = rightRoot[:len(rightRoot)-1]
				result = append(result, curs.Val)
				if curs.Right != nil {
					stack = append(stack, curs.Right)
					break
				}
			}
		}
	}
	return result
}
