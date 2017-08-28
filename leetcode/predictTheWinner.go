package leetcode

func PredictTheWinner(nums []int) bool {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}
	for length := 2 ; length <= len(nums) ; length++ {
		start, end := 0, 0
		for end < len(nums) - 1 {
			end = start + length - 1
			dp[start][end] = max(nums[start]-dp[start+1][end], nums[end]-dp[start][end-1])
			start += 1
		}
	}

	return dp[0][len(nums)-1] >= 0
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}