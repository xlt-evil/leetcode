package main

import (
	"fmt"
)

//解开锁
func openLock(deadends []string, target string) int {
	step := 0
	record := map[string]bool{}
	deads := map[string]bool{}
	for i, _ := range deadends {
		deads[deadends[i]] = true
	}
	Q := []string{"0000"}
	for len(Q) > 0 {
		length := len(Q) //当前步的队列
		for i := 0; i < length; i++ {
			cur := Q[0]
			Q = Q[1:]
			if cur == target {
				return step
			}
			_, Ok := deads[cur]
			if Ok {
				continue
			}
			for i, _ := range cur {
				temp := cur[i]
				// 每个数都可以向上转或向下移
				add, sub := temp+1, temp-1
				if temp == 57 {
					add = 48
				} else if temp == 48 {
					sub = 57
				}
				s1 := cur[:i] + string(add) + cur[i+1:]
				s2 := cur[:i] + string(sub) + cur[i+1:]
				_, ok := record[s1] // 判断是否已走过
				_, ok1 := deads[s1] //判断是死锁
				if !ok && !ok1 {
					record[s1] = true //记录判断锁
					Q = append(Q, s1) //添加至下一步的迭代中
				}
				_, ok = record[s2]
				_, ok1 = deads[s2]
				if !ok && !ok1 {
					record[s2] = true
					Q = append(Q, s2)
				}
			}
		}
		//当前可能的步数全部走完，步数加一
		step++
	}
	return -1
}

func main() {
	fmt.Println(openLock([]string{"0008"}, "0019"))
}
