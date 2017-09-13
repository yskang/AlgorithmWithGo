package leetcode

import (
	"math"
)

func MinimumTotal(triangle [][]int) int {
	return minimumTotal(triangle)
}

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, i+1)
	}

	dp[0][0] = triangle[0][0]

	for i := 1 ; i < len(triangle) ; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1 ; j < i ; j++ {
			dp[i][j] = minTriangle(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	minV := math.MaxInt64
	for _, v := range dp[len(dp)-1] {
		if minV > v {
			minV = v
		}
	}

	return minV
}

func minTriangle(a, b int) int {
	if a < b {
		return a
	}
	return b
}