package main

import "fmt"

func main() {
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)
	fmt.Println(findPeakGrid([][]int{
		{1, 4},
		{2, 3},
	}))
}

func findPeakGrid(mat [][]int) []int {
	res := []int{0, 0}
	m := mat[0][0]
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] > m {
				res[0] = i
				res[1] = j
				m = mat[i][j]
			}
		}
	}
	return res
}
