package main

import (
	"fmt"
)

func main() {
	fmt.Println(numRookCaptures([][]byte{
		{'.', '.','.','.','.','.','.','.'},
		{'.', '.','.','p','.','.','.','.'},
		{'.', '.','.','R','.','.','.','p'},
		{'.', '.','.','.','.','.','.','.'},
		{'.', '.','.','.','.','.','.','.'},
		{'.', '.','.','p','.','.','.','.'},
		{'.', '.','.','.','.','.','.','.'},
		{'.', '.','.','.','.','.','.','.'},
	}))
}

func numRookCaptures(board [][]byte) int {
	ans := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'R' {
				fmt.Println(i, j)
				upi := i
				// 上
				for upi >= 0 && board[upi][j] != 'B' {
					if board[upi][j] == 'p' {
						ans++
						break
					}
					upi--
				}
				// 下
				downi := i
				for downi <= 7 && board[downi][j] != 'B' {
					if board[downi][j] == 'p' {
						ans++
						break
					}
					downi++
				}
				// 左
				leftj := j
				for leftj >= 0 && board[i][leftj] != 'B' {
					if board[i][leftj] == 'p' {
						ans++
						break
					}
					leftj--
				}
				// 右
				rightj := j
				for rightj <= 7 && board[i][rightj] != 'B' {
					if board[i][rightj] == 'p' {
						ans++
						break
					}
					rightj++
				}
			}
		}
	}
	return ans
}