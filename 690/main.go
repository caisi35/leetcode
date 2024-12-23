package main

import (
	"fmt"
)

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func main() {
	emp1 := Employee{1, 5, []int{2, 3}}
	emp2 := Employee{2, 3, []int{}}
	emp3 := Employee{3, 3, []int{}}
	emps := []*Employee{&emp1, &emp2, &emp3}	
	fmt.Println(getImportence(emps, 1))	// 11
}

func getImportence(employees []*Employee, id int) int {
	ans := 0
	t := map[int]*Employee{}
	sub := []int{}
	for _, e := range employees {
		t[e.Id] = e
		if e.Id == id {
			sub = append(sub, e.Subordinates...)
			ans += e.Importance
		}
	}
	var dfs func(s []int)
	dfs = func(s []int) {
		for _, idx := range s {
			if v, ok := t[idx]; ok {
				ans += v.Importance
				if len(v.Subordinates) > 0 {
					dfs(v.Subordinates)
				}
			}
		}
	}
	dfs(sub)
	return ans
}
