package leetcode

import (
	"math"
)

func MaxProfit4(k int, prices []int) int {
	return maxProfit4(k, prices)
}

func maxProfit4(k int, prices []int) int {
	if k == 0 || len(prices) < 2 {
		return 0
	}

	if k > len(prices) {
		k = len(prices)
	}

	initEmpty := 0

	havesBefore := make([]int, k)
	haves := make([]int, k)
	emptiesBefore := make([]int, k)
	empties := make([]int, k)

	havesBefore[0] = -prices[0]

	for i := 1 ; i < len(prices) ; i++ {
		haves[0] = maxIV(havesBefore[0], initEmpty-prices[i])
		empties[0] = maxIV(emptiesBefore[0], havesBefore[0]+prices[i])
		havesBefore[0] = haves[0]
		emptiesBefore[0] = empties[0]
		for j := 1 ; j < k ; j++ {
			if i < 1+j {
				haves[j] = -1000
			} else {
				haves[j] = maxIV(havesBefore[j], emptiesBefore[j-1]-prices[i])
			}

			if i < 2+j {
				empties[j] = -1000
			} else {
				empties[j] = maxIV(emptiesBefore[j], havesBefore[j]+prices[i])
			}
			emptiesBefore[j] = empties[j]
			havesBefore[j] = haves[j]
		}
	}

	maxP := math.MinInt64
	for i := 0 ; i < k ; i++ {
		if maxP < empties[i] {
			maxP = empties[i]
		}
	}
	return maxP
}

func maxIV(a int, b int) int {
	if a > b {
		return a
	}
	return b
}