package main

import (
	"fmt"
)

func main() {
	ns1 := NeighborSum{[][]int{
		{3,6,0},
		{8,5,1},
		{2,4,7},
	}}
	fmt.Println(ns1.AdjacentSum(1))
	fmt.Println(ns1.AdjacentSum(3))
	fmt.Println(ns1.AdjacentSum(7))
	fmt.Println(ns1.AdjacentSum(2))
	fmt.Println(ns1.AdjacentSum(5))
	fmt.Println(ns1.AdjacentSum(8))
	fmt.Println(ns1.AdjacentSum(0))
	fmt.Println(ns1.DiagonalSum(1))
	fmt.Println(ns1.DiagonalSum(4))
	fmt.Println(ns1.DiagonalSum(8))
	fmt.Println(ns1.DiagonalSum(6))
	fmt.Println(ns1.DiagonalSum(3))
	fmt.Println(ns1.DiagonalSum(2))
	fmt.Println(ns1.DiagonalSum(6))
	fmt.Println(ns1.DiagonalSum(5))
	fmt.Println(ns1.DiagonalSum(7))
	fmt.Println(ns1.DiagonalSum(4))
	fmt.Println(ns1.DiagonalSum(0))

	ns2 := NeighborSum{[][]int{
		{1,2,0,3},
		{4,7,15, 6},
		{8,9,10,11},
		{12,13,14,5},
	}}
	fmt.Println(ns2.AdjacentSum(15))
	fmt.Println(ns2.DiagonalSum(9))

	ns := NeighborSum{[][]int{
		{0,1,2},
		{3,4,5},
		{6,7,8},
	}}
	fmt.Println(ns.AdjacentSum(1))
	fmt.Println(ns.AdjacentSum(4))
	fmt.Println(ns.DiagonalSum(4))
	fmt.Println(ns.DiagonalSum(8))
}

type NeighborSum struct {
	grid [][]int
}

func Constructor(grid [][]int) NeighborSum {
	return NeighborSum{grid}
}

func (ns *NeighborSum) AdjacentSum(value int) int {
	sum := 0
	for i := range ns.grid {
		for j := range ns.grid[i] {
			if ns.grid[i][j] == value {
				// 上
				if i > 0 {
					sum += ns.grid[i-1][j]
				}
				// 下
				if i < len(ns.grid)-1 {
					sum += ns.grid[i+1][j]
				}
				// 左
				if j > 0 {
					sum += ns.grid[i][j-1]
				}
				// 右
				if j < len(ns.grid[i])-1 {
					sum += ns.grid[i][j+1]
				}
			}
		}
	}
	return sum
}

func (ns *NeighborSum) DiagonalSum(value int) int {
	sum := 0
	for i := range ns.grid {
		for j := range ns.grid[i] {
			if ns.grid[i][j] == value {
				// 左上
				if i > 0 && j > 0 {
					sum += ns.grid[i-1][j-1]
				}
				// 右下
				if i < len(ns.grid)-1 && j < len(ns.grid[i])-1 {
					sum += ns.grid[i+1][j+1]
				}
				// 左下
				if j > 0 && i < len(ns.grid)-1 {
					sum += ns.grid[i+1][j-1]
				}
				// 右上
				if j < len(ns.grid[i])-1 && i > 0 {
					sum += ns.grid[i-1][j+1]
				}
			}
		}
	}
	return sum
}