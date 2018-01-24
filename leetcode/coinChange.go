package leetcode

import (
	"math"
)

func CoinChange(coins []int, amount int) int {
	return coinChange(coins, amount)
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	total := make([]int, amount+1)
	for _, coin := range coins {
		if coin >= len(total) {
			continue
		}
		total[coin] = 1
	}
	total[0] = 0
	for i := 1; i <= amount; i++ {
		total[i] = makeAmount(i, coins, total)
	}
	if total[amount] == math.MaxInt64 {
		return -1
	}
	return total[amount]
}

func makeAmount(number int, coins []int, total []int) int {
	minCount := math.MaxInt64
	// for i := 1; i <= number; i++ {
	// 	count := total[i] + total[number-i]
	// 	if count > 0 && count < minCount {
	// 		minCount = count
	// 	}
	// }

	for _, coin := range coins {
		if number-coin >= 0 && number-coin < len(total) && coin < len(total) {
			count := total[number-coin] + total[coin]
			if count > 0 && count < minCount {
				minCount = count
			}
		}
	}

	total[number] = minCount
	return minCount
}
