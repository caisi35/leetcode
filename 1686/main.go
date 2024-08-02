package main

import (
	"fmt"
	"sort"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
	t.ShowB()
	fmt.Println(stoneGameVI([]int{1, 3}, []int{2, 1}))
}

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	values := make([][]int, n)
	for i := 0; i < n; i++ {
		values[i] = []int{aliceValues[i] + bobValues[i], aliceValues[i], bobValues[i]}
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i][0] > values[j][0]
	})
	aliceSum, bobSum := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			aliceSum += values[i][1]
		} else {
			bobSum += values[i][2]
		}
	}
	if aliceSum > bobSum {
		return 1
	} else if aliceSum == bobSum {
		return 0
	} else {
		return -1
	}
}
