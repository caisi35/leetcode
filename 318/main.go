package main

import (
	"fmt"
	// "sort"
)

func main() {
	var ans float64 = 15 + 25 + 5.2 + 5
	fmt.Println(ans)
	fmt.Println(maxProduct([]string{"eecdc","fdf","bfeccdffd","eaaef","eeea","cfa","eafdacddc","dcebf","bbdef","ccaede","df","caede","eadcdaeaec","eeccbdac","eabbcebadbb","feaeebbccc","dadcbf","fcfcfefc","dfdefeefadf","cf","adbebcaded","adbfa","bfbaa","eea","eea","ccee","ebacfcfb","cd","ddeb","aabcc","edadec","cbc","dbcbe","bcbaff","aedeabcf","cfaddfe","dbe","badd","ddccfcfee","ebcfdaecdac","efcfeefdb","ea","afbe","dcad","ddae","fccff","bebcacdb","ec","defcb","fdcecc","bc","bfbcbf","fff","beddb","cadccaf","fadbcaffc","cbcabecfc","fdadcbefde","efbcaad","abffafef","dacbea","feade","eadcfdcbffc","fed","cddcde","adbdba","ddcbbfabafd","dcbd","cddceacca","cbecbbdfde","ecefa","fdbbcb","adaf","ae","abbabeddc","aedbb","bbffbcfa","bddcfcdddc","acfacaaefd","becbaecff","eabceedaccb","caffcadbd","cdccc","ebdaeadc","cddededf","fefeeafd","eab","bfbddeea","edbccc","cfeecd","feaacaafbda","bcaefafdc","fadeebaa","efeacfbc","bb","dbabf","dcebfbdf","baeecbd","abcb","ffdbb","ddebfc","edfadc","eaeffffa","aebdc","cdabc","dacddeedad","fae","dbebdcadeec"}))
}

func maxProduct_(words []string) int {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}
	ans := 0
	for i, x := range masks {
		for j, y := range masks[:i] {
			if x & y == 0 && len(words[i]) * len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return ans
}

func maxProduct(words []string) int {
	n := len(words)
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			for _, w1 := range words[i] {
				for _, w2 := range words[j] {
					if w1 == w2 {
						goto endFor
					}
				}
			}
			ans = max(ans, len(words[i])*len(words[j]))
		endFor:
		}
	}

	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
