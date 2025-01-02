package main

import (
	"fmt"
)



func main() {
	obj := Constructor()
	// fmt.Println(obj.Book(47, 50))
	// fmt.Println(obj.Book(33, 41))
	// fmt.Println(obj.Book(39, 45))
	// fmt.Println(obj.Book(33, 42))

	// fmt.Println(obj.Book(19, 25))
	// fmt.Println(obj.Book(18, 27))	// false

	fmt.Println(obj.Book(37, 50))
	fmt.Println(obj.Book(33, 50))
	fmt.Println(obj.Book(60, 70))
	fmt.Println(obj.Book(10, 20))
}

type MyCalendar struct {
    l [][]int
}


func Constructor() MyCalendar {
    return MyCalendar{l: [][]int{}}
}


func (this *MyCalendar) Book(startTime int, endTime int) bool {
    for _, v := range this.l {
		if endTime > v[0] && startTime < v[1] {
			return false
		}
	}
	this.l = append(this.l, []int{startTime, endTime})
	return true
}


