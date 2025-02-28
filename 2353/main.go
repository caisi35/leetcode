package main

import (
	"container/heap"
	"fmt"
)

func main() {
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	cuisines := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	ratings := []int{9, 12, 8, 15, 14, 7}

	c := Constructor(foods, cuisines, ratings)
	fmt.Println(c.HighestRated("korean")) // kimchi
	fmt.Println(c.HighestRated("japanese")) // ramen
	c.ChangeRating("sushi", 16)
	fmt.Println(c.HighestRated("japanese")) // sushi
	c.ChangeRating("ramen", 16)
	fmt.Println(c.HighestRated("japanese")) // ramen


}

type FoodRating struct {
	rating int
	food   string
}

type FoodHeap []FoodRating

func (h FoodHeap) Len() int { return len(h) }
func (h FoodHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].food < h[j].food
	}
	return h[i].rating < h[j].rating
}
func (h FoodHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *FoodHeap) Push(x interface{}) {
	*h = append(*h, x.(FoodRating))
}
func (h *FoodHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type FoodRatings struct {
	foodMap map[string]struct {
		rating  int
		cuisine string
	}
	ratingMap map[string]*FoodHeap
	n         int
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	n := len(foods)
	foodMap := make(map[string]struct {
		rating  int
		cuisine string
	})
	ratingMap := make(map[string]*FoodHeap)

	for i := 0; i < n; i++ {
		food := foods[i]
		cuisine := cuisines[i]
		rating := ratings[i]
		foodMap[food] = struct {
			rating  int
			cuisine string
		}{
			cuisine: cuisine,
			rating:  rating,
		}

		if ratingMap[cuisine] == nil {
			ratingMap[cuisine] = &FoodHeap{}
		}

		heap.Push(ratingMap[cuisine], FoodRating{n - rating, food})
	}
	return FoodRatings{foodMap: foodMap, ratingMap: ratingMap, n: n}
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	entry := fr.foodMap[food]
	cuisine := entry.cuisine
	heap.Push(fr.ratingMap[cuisine], FoodRating{fr.n - newRating, food})
	entry.rating = newRating
	fr.foodMap[food] = entry
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	h := fr.ratingMap[cuisine]
	for h.Len() > 0 {
		top := (*h)[0]
		if fr.foodMap[top.food].rating == fr.n-top.rating {
			return top.food
		}
		heap.Pop(h)
	}
	return ""
}
