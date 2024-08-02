package main

import (
	"fmt"
)

type Test struct {
	name string
}

func (t *Test) Point(){
	fmt.Println(t.name)
}

func main() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key,val := range slice {
		v := val
		m[key] = &v
	}

	for k,v := range m {
		fmt.Println(k,"->",*v)  
		// 0 3
		// 1 3
		// 2 3
		// 3 3
	}

	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}

	for _,t := range ts {
		//fmt.Println(reflect.TypeOf(t))
		_t := t
		defer _t.Point()
	}

	fmt.Println(maxTaxiEarnings(5, [][]int{{2, 5, 4}, {1, 5, 1}}))
}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	ans := int64(0)
	return ans
}