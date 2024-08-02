package main

import "fmt"

func theMaximumAchievableX(num int, t int) int {
    return num + t * 2
}

func main() {
	fmt.Println(theMaximumAchievableX(4, 1))
}