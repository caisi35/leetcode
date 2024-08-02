package main

import "fmt"

func GetValue() int {
	return 1
}

func main() {
	// i := GetValue()
	// switch i.(type) {
	// case int:
	// 	println("int")
	// case string:
	// 	println("string")
	// case interface{}:
	// 	println("interface")
	// default:
	// 	println("unknown")
	// }

	fmt.Println(countWords([]string{"leetcode","is","amazing","as","is"}, []string{"amazing","leetcode","is"}))
}

func countWords(words1 []string, words2 []string) int {
	ans := 0 
	m := map[string]int{}
	m2 := map[string]int{}
	for _, i := range words1 {
		m[i] += 1
	}
	for _, j := range words2 {
		m2[j] += 1
	}
	for k, v := range m {
		if v2, ok := m2[k]; ok && v == 1 && v2 == 1 {
			ans++
		}
	}
	return ans
}