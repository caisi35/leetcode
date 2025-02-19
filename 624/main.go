package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDistance([][]int{{1}, {2}}))	// 1
	fmt.Println(maxDistance([][]int{{1,5}, {3, 4}}))	// 3
	fmt.Println(maxDistance([][]int{{1,2,3}, {4,5}, {1,2,3}}))	// 4
	fmt.Println(maxDistance([][]int{{1}, {1}}))	// 0
}

func maxDistance(arrays [][]int) int {
	mx, mi := arrays[0][len(arrays[0])-1], arrays[0][0]
	d := 0
	for i := 1; i < len(arrays); i++ {
		cmi := arrays[i][0]
		cmx := arrays[i][len(arrays[i])-1]

		d1 := abs(mx - cmi)
		d2 := abs(cmx - mi)

		if d1 > d {
			d = d1
		}
		if d2 > d {
			d = d2
		}
		if cmi < mi {
			mi = cmi
		}
		if cmx > mx {
			mx = cmx
		}
	}
	return d
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}