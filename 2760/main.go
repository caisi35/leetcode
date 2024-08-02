package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
	fmt.Println(longestAlternatingSubarray([]int{2, 7}, 9))

	fmt.Println(longestAlternatingSubarray([]int{3,2,5,4}, 5))
}

func longestAlternatingSubarray(nums []int, threshold int) int {
	ans := 0
	for i, v := range nums {
		if v % 2 == 0 && v <= threshold {
			m := 1
			for k := i+1; k < len(nums); k++ {
				if nums[k] <= threshold && nums[k-1] % 2 != nums[k] %2 {
					m += 1
				} else {
					break
				}
			}
			if m > ans {
				ans = m
			}
		}
	}
	return ans
}