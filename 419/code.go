package main

import "fmt"

func countBattleships(board [][]byte) (ans int) {
	for i, row := range board {
		for j, ch := range row {
			if ch == 'X' {
				// 纵横修改“X”为“。”
				row[j] = '.' //当前的修改
				// 横向的修改
				for k := j + 1; k < len(board[0]); k++ {
					if row[k] != 'X' {
						break
					}
					row[k] = '.'
				}
				// 纵向修改
				for k := i + 1; k < len(board); k++ {
					if board[k][j] != 'X' {
						break
					}
					board[k][j] = '.'
				}
				ans++
			}
		}
	}
	return
}

func main() {
	input := [][]byte {{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}
	ans := countBattleships(input)
	fmt.Println(ans)
}
