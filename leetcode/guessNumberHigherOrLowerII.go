package leetcode

func GetMoneyAmount(n int) int {
	return getMoneyAmount(n)
}

func getMoneyAmount(n int) int {

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for l := 0 ; l < n ; l++ {
		for s := 0 ; s+l < n ; s++ {
			if l == 0 {
				dp[s][s+l] = 0
			} else if l == 1 {
				dp[s][s+1] = s
			} else {
				for k := s + 1 ; k < s+l; k++ {
					dp[s][s+l] = minMon(dp[s][s+l], k + maxMon(dp[s][k-1], dp[k+1][s+l]))
				}
			}
		}
	}

	return dp[0][n-1]
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
