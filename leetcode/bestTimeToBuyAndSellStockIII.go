package leetcode


func MaxProfit3(prices []int) int {
	return maxProfit3(prices)
}

func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	firstFree := make([]int, len(prices))
	firstHave := make([]int, len(prices))
	secondFree := make([]int, len(prices))
	secondHave := make([]int, len(prices))
	lastFree := make([]int, len(prices))

	firstFree[0] = 0
	firstHave[0] = -prices[0]
	secondFree[0] = 0

	for i := 1 ; i < len(prices) ; i++ {
		firstFree[i] = firstFree[i-1]
		firstHave[i] = maxIII(firstHave[i-1], firstFree[i-1] - prices[i])
		secondFree[i] = maxIII(secondFree[i-1], firstHave[i-1] + prices[i])
		if i < 2 {
			secondHave[i] = -10000
		} else {
			secondHave[i] = maxIII(secondHave[i-1], secondFree[i-1] - prices[i])
		}
		if i < 3 {
			lastFree[i] = -10000
		} else {

			lastFree[i] = maxIII(lastFree[i-1], secondHave[i-1] + prices[i])
		}
	}

	return maxIII(secondFree[len(prices)-1], lastFree[len(prices)-1])
}

func maxIII(a int, b int) int {
	if a > b {
		return a
	}
	return b
}