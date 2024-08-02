package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	positive_feedback := []string{"smart","brilliant","studious"}
	negative_feedback := []string{"not"}
	report := []string{"this student is studious","the student is smart"}
	student_id := []int{1, 2}
	k := 2
	fmt.Println(topStudents(positive_feedback, negative_feedback, report, student_id, k))
	// 输出：[1,2]
}

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	ans := make([]int, k)
	p_map := make(map[string]int, len(positive_feedback))
	for i, v := range positive_feedback {
		p_map[v] = i
	}
	n_map := make(map[string]int, len(negative_feedback))
	for i, v := range negative_feedback {
		n_map[v] = i
	}
	// fmt.Println(p_map, n_map)
	slice_report := make([][]string, len(report))
	for i := range slice_report {
		slice_report[i] = strings.Split(report[i], " ")
	}
	// fmt.Println(slice_report[0])

	ans_map := make(map[int]int, len(student_id))
	for _, s := range student_id {
		ans_map[s] = 0
	}
	// fmt.Println(ans_map)

	for i, s_report := range slice_report {
		std_id := student_id[i]
		for _, word := range s_report {
			if _, ok := p_map[word]; ok {
				ans_map[std_id] += 3
			} else if _, ok := n_map[word]; ok {
				ans_map[std_id] -= 1
			}
		}
	}

	// fmt.Println(ans_map)

	ans_slice := [][]int{}
	for k, v := range ans_map {
		ans_slice = append(ans_slice, []int{k, v})
	}

	sort.Slice(ans_slice, func(i, j int) bool {
		if ans_slice[i][1] == ans_slice[j][1] {
			return ans_slice[i][0] < ans_slice[j][0]
		}
		return ans_slice[i][1] > ans_slice[j][1]
	})

	// fmt.Println(ans_slice)
	for i, a := range ans_slice {
		ans[i] = a[0]
		if k-1 == i {
			break
		}
	}

	return ans
}


