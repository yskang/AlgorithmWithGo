package leetcode

import (
	"math"
)

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	free := make([]int, len(prices))
	have := make([]int, len(prices))
	sell := make([]int, len(prices))

	free[0] = 0
	have[0] = -prices[0]
	sell[0] = math.MinInt64

	for i := 1 ; i < len(prices) ; i++ {
		free[i] = max2(free[i-1], sell[i-1])
		have[i] = max2(have[i-1], free[i-1] - prices[i])
		sell[i] = have[i-1] + prices[i]
	}

	return max2(free[len(prices)-1], sell[len(prices)-1])
}

func max2(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

