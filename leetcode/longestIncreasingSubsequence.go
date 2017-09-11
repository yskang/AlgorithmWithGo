package leetcode

import "fmt"

func LengthOfLIS(nums []int) int {
	return lengthOfLIS(nums)
}

func lengthOfLIS(nums []int) int {
	fmt.Println(nums)
	dp := make([][]minMax, len(nums))
	for i := range nums {
		dp[i] = make([]minMax, len(nums))
	}

	for s := 0 ; s < len(nums) ; s++ {
		for l := 0 ; s+l < len(nums) ; l++ {
			dp[s][s+l] = getLIS(s, s+l, nums, dp)
		}
	}
	fmt.Println(dp)
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

	if dp[s+1][e].min > nums[s] && dp[s][e-1].max < nums[e] {
		return minMax{nums[s], nums[e], dp[s+1][e-1].length+2}
	} else if dp[s+1][e].min > nums[s] {
		return minMax{nums[s], dp[s+1][e].max, dp[s+1][e].length+1}
	} else if dp[s][e-1].max < nums[e] {
		return minMax{dp[s][e-1].min, nums[e], dp[s][e-1].length+1}
	}

	mininum := 0
	maximum := 0

	if nums[s] > dp[s+1][e].max {
		maximum = nums[s]
	} else {
		maximum = dp[s+1][e].max
	}

	if nums[e] < dp[s][e-1].min {
		mininum = nums[e]
	} else {
		mininum = dp[s][e-1].min
	}

	return minMax{mininum, maximum, dp[s+1][e-1].length}
}

type minMax struct {
	min int
	max int
	length int
}