package main

import (
	"fmt"
	"leecode/dynamicProgramming/question"
)

func main() {
	fmt.Println("动态规划1.最优子结构。2.重叠子问题。3.无后效性")
	fmt.Println("1.判断有多少种走法leetcode-63")
	fmt.Println(question.DP1([][]int{
		[]int{0, 0},
		[]int{0, 0},
	}))
	fmt.Println("2.判断到终点的最小路径leetcode-64")
	fmt.Println(question.MinPathSum([][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}))
	fmt.Println("3.三角形最短路径leetcode-120")
	fmt.Println(question.MinimumTotal([][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}))
	fmt.Println("4.房子小偷leetcode-198")
	fmt.Println(question.Rob([]int{10, 2, 1, 11}))
	fmt.Println("5.爬楼梯leecode-70")
	fmt.Println(question.ClimbStairs(3))
	fmt.Println("6.最小代价进入楼顶,leecode-746") // 一般求最大最小的dp需要保存两个值，当前选择最优结果，但是因为不知道后面的变化所以还要保存另一个状态的结果，然后到了最后在判断连个结果
	fmt.Println(question.MinCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println("7.买卖股票的最佳时机，leetcode-121")
	fmt.Println(question.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("8.区域和检索 - 数组不可变,leetcode-303")
	num := question.Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(num.SumRange(0, 2))
	fmt.Println(num.SumRange(2, 5))
	fmt.Println("9.最大子串,leetcode-53") //dp的空间一般可以是参数空间的一半甚至更少，经验之谈 = =
	fmt.Println(question.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println("10.判断谁能赢leetcode-486")
	fmt.Println(question.PredictTheWinner([]int{1, 5, 23, 7}))
	fmt.Println("11.买卖股票的最佳时机3,leetcode-121")
	fmt.Println(question.MaxProfit3([]int{1, 4, 0, 3}))
}
