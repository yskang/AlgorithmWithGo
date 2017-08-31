package leetcode


func FindTargetSumWays(nums []int, S int) int {
	return findTargetSumWays(nums, S)
}


func findTargetSumWays(nums []int, S int) int {
	count := 0
	findCase(nums, S, 0, 0, &count)

	return count
}

func findCase(nums []int, S int, i int, sum int, count *int) {
	if i == len(nums) {
		if sum == S {
			*count += 1
		}
	} else {
		findCase(nums, S, i+1, sum + nums[i], count)
		findCase(nums, S, i+1, sum - nums[i], count)
	}
}

func FindTargetSumWaysWithDP(nums []int, S int) int {
	return findTargetSumWaysWithDP(nums, S)
}

func findTargetSumWaysWithDP(nums []int, S int) int {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2001)
	}

	dp[0][nums[0]+1000] = 1
	dp[0][-nums[0]+1000] += 1

	for i, num := range nums[1:] {
		for sum := -1000; sum <= 1000; sum++ {
			if dp[i][sum+1000] > 0 {
				dp[i+1][sum+num+1000] += dp[i][sum+1000]
				dp[i+1][sum-num+1000] += dp[i][sum+1000]
			}
		}
	}

	if S > 1000 {
		return 0
	}
	return dp[len(nums)-1][S+1000]
}
