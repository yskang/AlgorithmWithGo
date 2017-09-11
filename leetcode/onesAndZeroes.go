package leetcode

import "strings"

func FindMaxForm(strs []string, m int, n int) int {
	return findMaxForm(strs, m, n)
}

func findMaxForm(strs []string, m int, n int) int {

	dp := make([][][]int, len(strs)+1)

	for i := 0 ; i <= len(strs) ; i++ {
		numZero, numOne := 0, 0
		if i > 0 {
			numZero = strings.Count(strs[i-1], "0")
			numOne = strings.Count(strs[i-1], "1")
		}
		dp[i] = make([][]int, m+1)

		for j := 0 ; j <= m ; j++ {
			dp[i][j] = make([]int, n+1)
			for k := 0 ; k <= n ; k++ {
				if i == 0 {
					dp[i][j][k] = 0
				} else if j >= numZero && k >= numOne {
					dp[i][j][k] = maxNums(dp[i-1][j][k], dp[i-1][j-numZero][k-numOne] + 1)
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}

func maxNums(a int, b int) int {
	if a > b {
		return a
	}
	return b
}