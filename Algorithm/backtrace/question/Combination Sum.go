package question

import "fmt"

//回溯法
//condidates 选择元素集
//result 结果集
//r 当前的部分解集
//start 限制选择元素的范围
//target//目标
func backtrack(condidates []int, result *[][]int, r *[]int, start int, target int) {
	if target == 0 { //当目标等于0时找到答案
		cur := make([]int, len(*r))
		copy(cur, *r)                  //值拷贝
		*result = append(*result, cur) //保存找到一次结果的答案
		return
	} else if target < 0 { //说明当前结果集不服和答案
		return
	}
	for i := start; i < len(condidates); i++ { // 选择
		if condidates[i] > target {
			continue
		}
		// 找到当前符合的元素
		*r = append(*r, condidates[i])                            //添加到当前的结果集
		backtrack(condidates, result, r, i, target-condidates[i]) //去寻找以该元素为基础的答案
		//去掉该元素，判断当前是否还有其他元素可以使用
		*r = (*r)[:len(*r)-1] //因为优先级关系所以加上括号
	}
}

// 找打满足target的所有解
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	r := make([]int, 0)
	backtrack(candidates, &result, &r, 0, target)
	return result
}

func main() {
	fmt.Println(combinationSum([]int{1, 7}, 7))
}
