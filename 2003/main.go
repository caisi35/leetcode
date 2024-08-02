package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	n := len(parents)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	var dfs func(int) (map[int]bool, int)
	dfs = func(node int) (map[int]bool, int) {
		geneSet := map[int]bool{nums[node]: true}
		for _, child := range children[node] {
			childGeneSet, y := dfs(child)
			res[node] = max(res[node], y)
			if len(childGeneSet) > len(geneSet) {
				geneSet, childGeneSet = childGeneSet, geneSet
			}
			for gene, _ := range childGeneSet {
				geneSet[gene] = true
			}
		}
		for geneSet[res[node]] {
			res[node]++
		}
		return geneSet, res[node]
	}
	dfs(0)
	return res
}

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1. 
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)  

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
	fmt.Println(smallestMissingValueSubtree([]int{-1, 0, 0, 2}, []int{1, 2, 3, 4}))
}
