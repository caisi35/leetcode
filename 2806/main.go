package main

import "fmt"

func main() {
	fmt.Println(accountBalanceAfterPurchase(9))
	fmt.Println(accountBalanceAfterPurchase(15))
	fmt.Println(accountBalanceAfterPurchase(14))
}

func accountBalanceAfterPurchase(purchaseAmount int) int {
    m := purchaseAmount % 10
	s := 0
	if m >= 5 {
		s = 10
	}

	return 100 - (purchaseAmount - m + s)
}