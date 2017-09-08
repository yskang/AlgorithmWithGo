package leetcode

func MaxProfit2(prices []int) int {
	return maxProfit2(prices)
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	free := make([]int, len(prices))
	have := make([]int, len(prices))

	free[0] = 0
	have[0] = -prices[0]

	for i := 1 ; i < len(prices) ; i++ {
		free[i] = max3(free[i-1], have[i-1]+prices[i])
		have[i] = max3(have[i-1], free[i-1]-prices[i])
	}

	return free[len(prices)-1]
}

func max3(a int, b int) int {
	if a > b {
		return a
	}
	return b
}