package main

import "fmt"

type ATM struct {
	value []int64
	cnt   []int64
}

func Constructor() ATM {
	return ATM{cnt: make([]int64, 5), value: []int64{20, 50, 100, 200, 500}}
}

func (atm *ATM) Deposit(banknotesCount []int) {
	for i := range banknotesCount {
		atm.cnt[i] += int64(banknotesCount[i])
	}
}

func (atm *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	for i := 4; i >= 0; i-- {
		res[i] = int(min(atm.cnt[i], int64(amount)/atm.value[i]))
		amount -= res[i] * int(atm.value[i])
	}
	if amount > 0 {
		return []int{-1}
	}
	for i := 0; i < 5; i++ {
		atm.cnt[i] -= int64(res[i])
	}
	return res
}

func main() {
	obj := Constructor()
	obj.Deposit([]int{0,0,1,2,1})
	fmt.Println(obj.Withdraw(600))
	obj.Deposit([]int{0,1,0,1,1})
	fmt.Println(obj.Withdraw(600))
}