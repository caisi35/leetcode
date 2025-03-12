package main

import (
	"fmt"
)

func main() {
	fmt.Println(countOfSubstrings("aeioubaabca", 1))
}

func countOfSubstrings(word string, k int) int {
	vowels := map[byte]bool {
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	count := func(m int) int {
		n := len(word)
		res := 0
		consonants := 0
		occur := make(map[byte]int)
		for i, j := 0, 0; i < n; i++ {
			for j < n && (consonants < m || len(occur) < 5) {
				if vowels[word[j]] {
					occur[word[j]]++
				} else {
					consonants++
				}
				j++
			}
			if consonants >= m && len(occur) == 5 {
				res += n - j + 1
			}
			if vowels[word[i]] {
				occur[word[i]]--
				if occur[word[i]] == 0 {
					delete(occur, word[i])
				}
			} else {
				consonants--
			}
		}
		return res
	}
	return count(k) - count(k+1)
}

func countOfSubstrings2(word string, k int) int {
	vowels := map[byte]bool {
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	n := len(word)
	count := 0
	left := 0
	right := 0
	vowelCount := make(map[byte]int)
	consonantCount := 0
	for right < n {
		char := word[right]
		if vowels[char] {
			vowelCount[char]++
		} else {
			consonantCount++
		}
		for consonantCount > k {
			leftChar := word[left]
			if vowels[leftChar] {
				vowelCount[leftChar]--
				if vowelCount[leftChar] == 0 {
					delete(vowelCount, leftChar)
				}
			} else {
				consonantCount--
			}
			left++
		}
		if consonantCount == k && len(vowelCount) == 5 {
			count++
		}
		right++
	}
	return count
}