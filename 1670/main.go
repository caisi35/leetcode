package main

import (
	"container/list"
	"fmt"
)

func change(s ...int) {
	fmt.Println(len(s), cap(s), s)
	s = append(s, 3)
	fmt.Println(len(s), cap(s), s)

}

func main() {
	// slice := make([]int, 5, 5)
	// slice[0] = 1
	// slice[1] = 2
	// change(slice...)
	// fmt.Println(slice) // 1 2 0 0 0
	// change(slice[0:2]...)
	// fmt.Println(slice) // 1 2 3 0 0
	// fmt.Println(5 / 2)

	obj := Constructor()
	obj.PushFront(1)
	obj.PushBack(2)
	obj.PushMiddle(3)
	obj.PushMiddle(4)
	fmt.Println(obj)
	param_4 := obj.PopFront()
	param_5 := obj.PopMiddle()
	param_55 := obj.PopMiddle()
	param_6 := obj.PopBack()
	param_7 := obj.PopFront()
	fmt.Println(param_4, param_5, param_55, param_6, param_7)
}

type FrontMiddleBackQueue struct {
	left *list.List
	right *list.List
}


func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		left: list.New(),
		right: list.New(),
	}
}


func (this *FrontMiddleBackQueue) PushFront(val int)  {
	this.left.PushFront(val)
	if this.left.Len() == this.right.Len() + 2 {
		this.right.PushFront(this.left.Back().Value.(int))
		this.left.Remove(this.left.Back())
	}
}


func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
	if this.left.Len() == this.right.Len() + 1 {
		this.right.PushFront(this.left.Back().Value.(int))
		this.left.Remove(this.left.Back())
	}
	this.left.PushBack(val)
}


func (this *FrontMiddleBackQueue) PushBack(val int)  {
	this.right.PushBack(val)
	if this.left.Len() + 1 == this.right.Len() {
		this.left.PushBack(this.right.Front().Value.(int))
		this.right.Remove(this.right.Front())
	}
}


func (this *FrontMiddleBackQueue) PopFront() int {
	if this.left.Len() == 0 {
		return -1
	}
	val := this.left.Front().Value.(int)
	this.left.Remove(this.left.Front())
	if this.left.Len() + 1 == this.right.Len() {
		 this.left.PushBack(this.right.Front().Value.(int))
		 this.right.Remove(this.right.Front())
	}
	return val
}


func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.left.Len() == 0 {
		return -1
	}
	val := this.left.Back().Value.(int)
	this.left.Remove(this.left.Back())
	if this.left.Len() + 1 == this.right.Len() {
		this.left.PushBack(this.right.Front().Value.(int))
		this.right.Remove(this.right.Front())
	}
	return val
}


func (this *FrontMiddleBackQueue) PopBack() int {
	if this.left.Len() == 0 {
		return -1
	}
	if this.right.Len() == 0 {
		val := this.left.Back().Value.(int)
		this.left.Remove(this.left.Back())
		return val
	} else {
		val := this.right.Back().Value.(int)
		this.right.Remove(this.right.Back())
		if this.left.Len() == this.right.Len() + 2 {
			this.right.PushFront(this.left.Back().Value.(int))
			this.left.Remove(this.left.Back())
		}
		return val
	}
}


// type FrontMiddleBackQueue struct {
// 	queue []int
// }

// func Constructor() FrontMiddleBackQueue {
// 	queue := FrontMiddleBackQueue{queue: []int{}}
// 	return queue
// }

// func (this *FrontMiddleBackQueue) PushFront(val int) {
// 	t := make([]int, len(this.queue)+1)
// 	t[0] = val
// 	t = append(t, this.queue...)
// 	fmt.Println("push front t:", t)
// 	this.queue = t
// 	fmt.Println("push front queue:", t)
// }

// func (this *FrontMiddleBackQueue) PushMiddle(val int) {
// 	l := len(this.queue)
// 	mid := 0
// 	if l%2 == 0 {
// 		mid = l / 2
// 	} else {
// 		mid = l / 2
// 	}
// 	t := make([]int, len(this.queue)+1)
// 	// t = append(t, this.queue[:mid]...)
// 	append_s(this.queue[:mid], t)
// 	t[mid] = val
// 	// t = append(t, this.queue[mid:]...)
// 	this.append_f(this.queue[mid:], t, mid)
// 	fmt.Println("push mid t:", t)
// 	this.queue = t
// 	fmt.Println("push mid queue:", t)
// }

// func append_s(s, d []int) {
// 	for i := range s {
// 		d[i] = s[i]
// 	}
// }

// func (fq *FrontMiddleBackQueue) append_f(s, d []int, mid int) {
// 	for i := range fq.queue[mid:] {
// 		// fmt.Println(i)
// 		d[mid+i+1] = s[i]
// 	}
// }

// func (fq *FrontMiddleBackQueue) append_fp(queue, t []int, mid int) {
// 	for i, v := range queue[mid:] {
// 		// fmt.Println(i)
// 		t[mid+i] = v
// 	}
// }

// func (fq *FrontMiddleBackQueue) append_fpt(queue, t []int, mid int) {
// 	for i := mid; i < len(t); i++ {
// 		// fmt.Println(i)
// 		t[mid] = queue[mid+1]
// 	}
// }

// func (this *FrontMiddleBackQueue) PushBack(val int) {
// 	this.queue = append(this.queue, val)
// 	fmt.Println("push back queue:", this.queue)
// }

// func (this *FrontMiddleBackQueue) PopFront() int {
// 	if len(this.queue) == 0 {
// 		return -1
// 	}
// 	v := this.queue[0]
// 	this.queue = this.queue[1:]
// 	return v
// }

// func (this *FrontMiddleBackQueue) PopMiddle() int {
// 	l := len(this.queue)
// 	if l == 0 {
// 		return -1
// 	}
// 	mid := 0
// 	t := make([]int, len(this.queue)-1)

// 	if l%2 == 0 {
// 		mid = l / 2 - 1
// 		append_s(this.queue[:mid], t)
// 		this.append_fpt(this.queue[mid:], t, mid)
// 	} else {
// 		mid = l / 2
// 		append_s(this.queue[:mid], t)
// 		this.append_fp(this.queue[mid:], t, mid)
// 	}
// 	ans := this.queue[mid]
// 	this.queue = t
// 	return ans
// }

// func (this *FrontMiddleBackQueue) PopBack() int {
// 	if len(this.queue) == 0 {
// 		return -1
// 	}
// 	ans := this.queue[len(this.queue)-1]
// 	this.queue = this.queue[:len(this.queue)-1]
// 	return ans
// }

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */


// from django.utils.translation import ugettext  已移除(改为gettext)
// import ipaddr	已移除(改为ipaddress)
// from django_mysql.model import JSONField	已移除(改为django.db.models的JSONField)
// NullBooleanField	已移除(改为BooleanField)
// url		已移除(改为path)
// admin组件必需开启授权中间件	(已注释组件&路由path)
// sync_view无视图		(已注释)
// swagger doc路由不兼容	(已注释)

// 添加配置:
// ENABLE_REDIS_CONN = True   # 测试是否开启redis连接
// DEFAULT_AUTO_FIELD = 'django.db.models.BigAutoField'   # TODO 默认主键类型

