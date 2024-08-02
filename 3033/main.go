package main

import "fmt"

func modifiedMatrix(matrix [][]int) [][]int {
	m := make([]int, len(matrix[0]))

	for i, mat  := range matrix {
		for j := range mat {
			if matrix[i][j] > m[j] {
				m[j] = matrix[i][j]
			}
		}
	}
	for i, mat  := range matrix {
		for j := range mat {
			if matrix[i][j] == -1 {
				matrix[i][j] = m[j]
			}
		}
	}
	return matrix
}

func main() {
	fmt.Println(modifiedMatrix([][]int{
		{1, 2, -1},
		{4, -1, 6},
		{7, 8, 9},
	}))
}