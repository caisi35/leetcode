package main

import (
	"fmt"
)

func main() {
	c := Constructor("leetcode.com")
	c.Visit("google.com")
	c.Visit("facebook.com")
	c.Visit("youtube.com")
	fmt.Println(c.Back(1))	// facebook
	fmt.Println(c.Back(1)) 		// google
	fmt.Println(c.Forward(1))	// facebook
	c.Visit("linkedin.com")
	fmt.Println(c.Forward(2))	// linkedin
	fmt.Println(c.Back(2))	// google
	fmt.Println(c.Back(7))	//	leetcode
}

type BrowserHistory struct {
	curr int
    page []string
}


func Constructor(homepage string) BrowserHistory {

    return BrowserHistory{page: []string{homepage}, curr: 0}
}


func (this *BrowserHistory) Visit(url string)  {
	for len(this.page) > this.curr + 1 {
		this.page = this.page[:this.curr + 1]
	}
	this.page = append(this.page, url)
	this.curr++
}


func (this *BrowserHistory) Back(steps int) string {
	this.curr = max(this.curr - steps, 0)
	return this.page[this.curr]
}


func (this *BrowserHistory) Forward(steps int) string {
    this.curr = min(this.curr + steps, len(this.page) - 1)
	return this.page[this.curr]
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */