package main

import "fmt"

type Students struct{
	Name string
}

func main() {
	list := make(map[string]Students)
	stu := Students{"Tom"}
	list["Student"] = stu
	// list["Student"].Name = "Alice"
	fmt.Println(list["Student"])
	stu2 := list["Student"]
	stu2.Name = "Alice"
	list["Student"] = stu2
	fmt.Println(list["Student"])
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0,1,0,1,0,1,0,1}, 3))
}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	total := 0
	n := len(customers)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}
	increase := 0
	for i := 0; i < minutes; i++ {
		increase += customers[i] * grumpy[i]
	}
	maxIncrease := increase
	for i := minutes; i < n; i++ {
		increase = increase - customers[i-minutes] * grumpy[i-minutes] + customers[i] * grumpy[i]
		maxIncrease = max(maxIncrease, increase)
	}
	return total + maxIncrease
}