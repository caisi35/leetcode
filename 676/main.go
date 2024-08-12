package main

import (
	"fmt"
)



/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
func main() {
	obj := Constructor()
	obj.BuildDict([]string{"hello","hallo","leetcode"})
	param_2 := obj.Search("hhllo")
	fmt.Println(param_2)	// true
}

type MagicDictionary struct {
	l []string
	// m map[string]bool
}


func Constructor() MagicDictionary {
	return MagicDictionary{}
}


func (this *MagicDictionary) BuildDict(dictionary []string)  {
	this.l = dictionary
}


func (this *MagicDictionary) Search(searchWord string) bool {
	for _, w := range this.l {
		c := 0

		if len(w) != len(searchWord) {
			continue
		}
		for i, ch := range searchWord {

			if byte(ch) != w[i] {
				c++
			}
		}
		if c == 1 {
			return true
		}
	}

	return false
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */