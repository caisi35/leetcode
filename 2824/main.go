package main

import (
	"fmt"
	"sort"
)

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
	fmt.Println(countPairs([]int{-1, 1, 2, 3, 1}, 2))
	fmt.Println(countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2))

}

func countPairs(nums []int, target int) int {
	ans := 0
	sort.Ints(nums)
	n := len(nums)
	// m := make(map[int]int)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if (nums[i]+nums[j]) < target {
				// if v, ok := m[i]; ok && v == j {
				// 	continue
				// }
				// m[i] = j
				ans++
			}
		}
	}
	return ans
}


// curl 'http://127.0.0.1:5555/api/subnets/delete/' -X POST -H "Content-Type:application/json" -d '{"os_id":"2022guizhou_duoaz_siyouyun_test","ct_user_id":"1000000715","subnet_uuid":"aaa10163-55ad-42a2-b366-0f18bfcf710c"}'

// curl 'http://127.0.0.1:5555/api/subnets/query/?os_id=2022guizhou_duoaz_siyouyun_test&ct_user_id=1000000715&vpc_id=5ae97731-f2d3-49a0-be39-bdfa95dad7ba&l2gw=true'