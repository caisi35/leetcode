package main

import (
	"fmt"
)

func main() {
	obj := Constructor(5)
	fmt.Println(obj.Insert(3, "ccccc")) // []
	fmt.Println(obj.Insert(1, "aaaaa")) // [aaaaa]
	fmt.Println(obj.Insert(2, "bbbbb")) // [bbbbb ccccc]
	fmt.Println(obj.Insert(5, "eeeee")) // []
	fmt.Println(obj.Insert(4, "ddddd")) // ["ddddd", "eeeee"]
}

type OrderedStream struct {
	ptr int
	l   []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{l: make([]string, n+1), ptr: 1}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	//如果流存储有 id = ptr 的 (id, value) 对，
	// 则找出从 id = ptr 开始的 最长 id 连续递增序列 ，
	// 并 按顺序 返回与这些 id 关联的值的列表。然后，将 ptr 更新为最后那个  id + 1
	this.l[idKey] = value
	if this.l[idKey] != "" && this.ptr == idKey {
		ans := []string{}
		idx := 0
		for _, v := range this.l[idKey:] {
			if v != "" {
				idx++
				ans = append(ans, v)
			} else {
				break
			}
		}
		if idx != 0 {
			this.ptr = idx + this.ptr
		}
		return ans
	}
	return []string{}
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
