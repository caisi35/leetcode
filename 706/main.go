package main

import (
	"container/list"
	"fmt"
)

const base = 769


type entry struct {
	key, value int
}
type MyHashMap struct {
	data []list.List
}


func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (m *MyHashMap) hash(key int) int {
	return key % base
}


func (this *MyHashMap) Put(key int, value int)  {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	this.data[h].PushBack(entry{key, value})
}


func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}


func (this *MyHashMap) Remove(key int)  {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			this.data[h].Remove(e)
		}
	}
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

 func main() {
	obj := Constructor();
	obj.Put(1,1);
	obj.Remove(1);
	param_2 := obj.Get(1);

	fmt.Println(param_2)
 }