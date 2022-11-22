package main

import (
	"fmt"
	"math"
)

const (
	downwardThreshold   = 0.25
	upwardThreshold     = 1.25
	effBalanceIncrement = 0.25
	maxEffBalance       = 32
)

var (
	balance    float64
	effBalance float64
)

func main() {
	var deposits []int
	deposits = []int{1, 1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1}
	for idx, deposit := range deposits {
		balance += float64(deposit)
		if balance-downwardThreshold < effBalance || effBalance+upwardThreshold < balance {
			effBalance = math.Min(balance-math.Mod(balance, effBalanceIncrement), maxEffBalance)
		}
		fmt.Println(idx, " balance :", balance, "effective balance :", effBalance)
	}
}
