package main

import "fmt"

func main() {
	fmt.Println(numberOfRightTriangles([][]int{
		{0, 1, 0},
		{0, 1, 1},
		{0, 1, 0},
	}))
}

func numberOfRightTriangles2(grid [][]int) int64 {
	ans := 0
	for i, row := range grid {
		for j := range row {
			row_s := 0
			col_s := 0
			if grid[i][j] == 1 {
				for ri := 0; ri < len(row); ri++ {

					if grid[i][ri] == 1 {
						row_s++
					}
				}
				for ci := 0; ci < len(grid); ci++ {
					if grid[ci][j] == 1 {
						col_s++
					}
				}
				ans += (row_s - 1) * (col_s - 1)
			}
		}
	}

	return int64(ans)
}

func numberOfRightTriangles(grid [][]int) int64 {
	ans := 0
	col := make([]int, len(grid[0]))
	for i, row := range grid {
		for j := range row {
			col[j] += grid[i][j]
		}
	}

	for i, row := range grid {
		row_1 := 0
		for _, v := range row {
			row_1 += v
		}
		for j := range row {
			if grid[i][j] == 1 {
				ans += (col[j] - 1) * (row_1 - 1)
			}
		}

	}

	return int64(ans)
}

