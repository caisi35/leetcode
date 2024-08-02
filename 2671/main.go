package main

import "fmt"

type Foo struct {
	bar string
}

func main() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])

	obj := Constructor()
	obj.Add(3)
	obj.Add(3)
	// fmt.Println(obj.HasFrequency(1))
	// obj.DeleteOne(3)
	fmt.Println(obj.HasFrequency(1))

}

type FrequencyTracker struct {
	m map[int]int // number : total
	f map[int]int // total: total
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{m: map[int]int{}, f: map[int]int{}}
}

func (this *FrequencyTracker) Add(number int) {
	this.m[number]++

	this.f[this.m[number]] += 1

	this.f[this.m[number]-1] -= 1
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if v, ok := this.m[number]; ok && v > 0 {
		this.m[number]--
		
		this.f[this.m[number]+1] -= 1

		this.f[this.m[number]] += 1
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	v, ok := this.f[frequency]
	return ok && v > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
