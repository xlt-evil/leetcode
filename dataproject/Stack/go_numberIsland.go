package main

import (
	"fmt"
)

func main() {
	s := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	fmt.Println(numIslands(s))
}

// 利用栈深度优先解决到岛屿数量问题 -- 递归是系统的隐式栈
type point struct {
	i, j int
}

// 定义四个方向 上、左、下、右
var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p point) at(grid [][]byte) (byte, bool) {
	if p.i < 0 || p.i >= len(grid) { // 上、下越界
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) { // 左、右越界
		return 0, false
	}
	return grid[p.i][p.j], true
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func numIslands(grid [][]byte) int {
	islands := 0
	steps := make([][]byte, len(grid))
	for i, _ := range steps {
		steps[i] = make([]byte, len(grid[i]))
	}
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if steps[i][j] != 0 {
				continue
			}
			if grid[i][j] == 49 {
				islands++
			} else {
				continue
			}
			start := point{i, j}        //当前岛屿起始点
			steps[start.i][start.j] = 1 // 记录探索起始点
			ThisIslandArea(start, grid, steps)
		}
	}
	return islands
}

func ThisIslandArea(p point, maps, record [][]byte) {
	for _, dir := range dirs { // 判断四个方向的可探索点
		next := p.add(dir)
		val, ok := next.at(maps)
		if !ok || val == 48 { // 越界or遇到水了
			continue
		}
		val, ok = next.at(record)
		if !ok || val != 0 { // 越界or探索过了
			continue
		}
		record[next.i][next.j] = 1 // 记录探索点
		ThisIslandArea(next, maps, record)
	}
}
