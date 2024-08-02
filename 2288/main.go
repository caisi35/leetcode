package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(discountPrices("there are $1 $2 and 5$ candies in the shop", 50)) //"there are $0.50 $1.00 and 5$ candies in the shop"
	fmt.Println(discountPrices("1 2 $3 4 $5 $6 7 8$ $9 $10$", 100))	//"1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"

	fmt.Println(discountPrices("$10$", 100))
	fmt.Println(discountPrices("ka3caz4837h6ada4 r1 $602", 9))	//"ka3caz4837h6ada4 r1 $547.82"
	fmt.Println(discountPrices("$602", 9))	//"$547.82"

}

func discountPrices(sentence string, discount int) string {
	s := strings.Split(sentence, " ")
	for i, word := range s {
		if word[0] == '$' {
			p, err := strconv.Atoi(string(word[1:]))
			if err == nil {
				s[i] = fmt.Sprintf("$%.2f", (1.0 - (float64(discount) / 100.0)) * float64(p))
			}
		}
	}
	return strings.Join(s, " ")
}