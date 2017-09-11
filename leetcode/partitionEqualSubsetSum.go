package leetcode

func CanPartition(nums []int) bool {
	return canPartition(nums)
}

func canPartition(nums []int) bool {
	sumOfall := 0
	for _, num := range nums {
		sumOfall += num
	}
	if sumOfall % 2 == 1 {
		return false
	}

	target := sumOfall / 2

	dp := make([][]bool, len(nums))
	for i := 0 ; i < len(nums) ; i++ {
		dp[i] = make([]bool, target+1)
	}

	for i := 0 ; i < len(nums) ; i++ {
		dp[i][0] = true
	}

	for j := 1 ; j <= target ; j++ {
		dp[0][j] = false
	}
	for j := 1 ; j <= target ; j++ {
		for i := 1 ; i < len(nums) ; i++ {
			if j - nums[i] >= 0 {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
			//fmt.Println("find dp [", i, j, "]: ", dp[i][j])
		}
	}
	//fmt.Println(dp)
	return dp[len(nums)-1][target]
}