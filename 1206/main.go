package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	s := Constructor()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Search(0)
	s.Add(4)
	s.Search(1)
	s.Erase(0)
	s.Erase(1)
	fmt.Println(s.Search(1))
}

const maxLevel = 32
const pFactor = 0.25

type SkiplistNode struct {
	val int
	forward []*SkiplistNode
}

type Skiplist struct {
	head *SkiplistNode
	level int
}

func Constructor() Skiplist {
	return Skiplist{&SkiplistNode{-1, make([]*SkiplistNode, maxLevel)}, 0}
}

func (Skiplist) randomLevel() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}

func (s *Skiplist) Search(target int) bool {
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].val < target {
			curr = curr.forward[i]
		}
	}
	curr = curr.forward[0]
	return curr != nil && curr.val == target
}

func (s *Skiplist) Add(num int) {
	update := make([]*SkiplistNode, maxLevel)
	for i := range update {
		update[i] = s.head
	}
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	lv := s.randomLevel()
	s.level = max(s.level, lv)
	newNode := &SkiplistNode{num, make([]*SkiplistNode, lv)}
	for i, node := range update[:lv] {
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}
}

func (s *Skiplist) Erase(num int) bool {
	update := make([]*SkiplistNode, maxLevel)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	curr = curr.forward[0]
	if curr == nil || curr.val != num {
		return false
	}
	for i := 0; i < s.level && update[i].forward[i] == curr; i++ {
		update[i].forward[i] = curr.forward[i]
	}
	for s.level > 1 && s.head.forward[s.level-1] == nil {
		s.level--
	}
	return true
}

