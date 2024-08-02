package main

import (
	"math"
	"sort"
	"fmt"
)

func findRadius(houses []int, heaters []int) int {
    ans := 0
    sort.Ints(heaters)
    for _, house := range houses {
        j := sort.SearchInts(heaters, house + 1)
        minDis := math.MaxInt32
        if j < len(heaters) {
            minDis = heaters[j] - house
        }
        i := j - 1
        if i >= 0 && house - heaters[i] < minDis {
            minDis = house - heaters[i]
        }
        if minDis > ans {
            ans = minDis
        }
    }
    return ans
}

func main() {
	houses := []int{1,2,3,4}
	heaters := []int{1,4}
	ans := findRadius(houses, heaters)
	fmt.Println("ans: ", ans)
}