package main

import "fmt"

type point struct {
	i int
	j int
}

var ps = []point{
	point{1, 0},
	point{0, -1},
	point{-1, 0},
	point{0, 1},
}

func IsOut(board [][]byte, p point) bool {
	length := len(board)
	if p.i >= length || p.i < 0 {
		return false
	}
	if p.j >= len(board[p.i]) || p.j < 0 {
		return false
	}
	return true
}

func (this point) Add(point point) point {
	this.i += point.i
	this.j += point.j
	return this
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if len(word) == FindSingeWord(point{i, j}, point{-1, -1}, word, 0, board) {
				return true
			}
		}
	}
	return false
}

func FindSingeWord(cur, last point, word string, index int, board [][]byte) int {
	length := 0
	if index == len(word) {
		return index
	}

	if IsOut(board, cur) && board[cur.i][cur.j] == word[index] {
		for i, _ := range ps {

			if cur.Add(ps[i]) == last {
				continue
			}
			length = FindSingeWord(cur.Add(ps[i]), cur, word, index+1, board)
			if length == len(word) {
				return length
			}
		}
	}
	return length
}

func main() {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "SFE"))
}
