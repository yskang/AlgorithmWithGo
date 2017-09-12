package leetcode

import (
	"math"
)

func GetMoneyAmount(n int) int {
	return getMoneyAmount(n)
}

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for l := 1 ; l <= n ; l++ {
		for s := 1 ; s+l <= n ; s++ {
			if l == 0 {
				dp[s][s+l] = 0
			} else if l == 1 {
				dp[s][s+l] = s
			} else {
				dp[s][s+l] = math.MaxInt16
				for k := s ; k <= s+l ; k++ {
					if s == k {
						dp[s][s+l] = minMon(dp[s][s+l], k + maxMon(0, dp[k+1][s+l]))
					} else if k == s+l {
						dp[s][s+l] = minMon(dp[s][s+l], k + maxMon(dp[s][k-1], 0))
					} else {
						dp[s][s+l] = minMon(dp[s][s+l], k + maxMon(dp[s][k-1], dp[k+1][s+l]))
					}
				}
			}
		}
	}

	return dp[1][n]
}

func minMon(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxMon(a, b int) int {
	if a > b {
		return a
	}
	return b
}
