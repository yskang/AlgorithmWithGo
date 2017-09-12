package leetcode

func LengthOfLIS(nums []int) int {
	return lengthOfLIS(nums)
}

func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	//fmt.Println(nums)
	dp := make([]int, len(nums))
	dp[0] = 1
	maxLIS := -1
	for i := 1 ; i < len(nums) ; i++ {
		maxDp := -1
		for j := i-1 ; j >= 0 ; j-- {
			if nums[j] < nums[i] {
				if maxDp < dp[j] {
					maxDp = dp[j]
				}
			}
		}
		if maxDp == -1 {
			dp[i] = 1
		} else {
			dp[i] = maxDp + 1
		}

		if dp[i] > maxLIS {
			maxLIS = dp[i]
		}
		//	fmt.Println(dp)
	}
	return maxLIS
}

// solution of Memory Limit Exceeded
func lengthOfLIS_MLE(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	dp := make([][]minMax, len(nums))
	for i := range nums {
		dp[i] = make([]minMax, len(nums))
	}

	for l := 0 ; l < len(nums) ; l++ {
		for s := 0 ; s+l < len(nums) ; s++ {
			dp[s][s+l] = getLIS(s, s+l, nums, dp)
		}
	}

	return dp[0][len(nums)-1].length
}

func getLIS(s int, e int, nums []int, dp [][]minMax) minMax {
	if s == e {
		return minMax{nums[s], nums[s], 1}
	} else if e - s == 1 {
		if nums[s] < nums[e] {
			return minMax{nums[s], nums[e], 2}
		}
		return minMax{nums[e], nums[s], 1}
	}

	if dp[s+1][e].min > nums[s] {
		return minMax{nums[s], dp[s+1][e].max, dp[s+1][e].length+1}
	} else if dp[s][e-1].max < nums[e] {
		return minMax{dp[s][e-1].min, nums[e], dp[s][e-1].length+1}
	}

	if dp[s+1][e].length > dp[s][e-1].length {
		return dp[s+1][e]
	}
	return dp[s][e-1]
}

type minMax struct {
	min int
	max int
	length int
}