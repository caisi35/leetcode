package main

import "fmt"

func main() {
	fmt.Println(pushDominoes("LLLL....RRRRLR"))
}

func pushDominoes(dominoes string) string {
	s := []byte(dominoes)
	i, n, left := 0, len(s), byte('L')
	for i < n {
		j := i
		for j < n && s[j] == '.' {
			j++
		}
		right := byte('R')
		if j < n {
			right = s[j]
		}
		if left == right {
			for i < j {
				s[i] = right
				i++
			}
		} else if left == 'R' && right == 'L' {
			k := j - 1
			for i < k {
				s[i] = 'R'
				s[k] = 'L'
				i++
				k--
			}
		}
		left = right
		i = j + 1
	}
	return string(s)
}