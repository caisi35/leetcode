package main

import "fmt"

type pair struct {
	start int
	end   int
}
type MyCalendarTwo struct {
	booked   []pair
	overlaps []pair
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(startTime int, endTime int) bool {
	for _, p := range this.overlaps {
		if p.start < endTime && p.end > startTime {
			return false
		}
	}
	for _, p := range this.booked {
		if p.start < endTime && p.end > startTime {
			this.overlaps = append(this.overlaps, pair{max(p.start, startTime), min(endTime, p.end)})
		}
	}
	this.booked = append(this.booked, pair{start: startTime, end: endTime})
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */

func main() {
	obj := Constructor()
	fmt.Println(obj.Book(10, 20))
	fmt.Println(obj.Book(40, 50))
	fmt.Println(obj.Book(60, 70))
	fmt.Println(obj.Book(20, 70))
	fmt.Println(obj.Book(10, 70))
	fmt.Println(obj.Book(40, 70))
}