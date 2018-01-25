package leetcode

import (
	"fmt"
	"math"
)

func MaxProduct(nums []int) int {
	return maxProduct(nums)
}

func maxProduct(nums []int) int {
	r, imax, imin := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		if nums[i] < imax*nums[i] {
			imax = imax * nums[i]
		} else {
			imax = nums[i]
		}
		if nums[i] < imin*nums[i] {
			imin = nums[i]
		} else {
			imin = imin * nums[i]
		}
		if r < imax {
			r = imax
		}
		fmt.Println("min, max", imin, imax)
	}
	return r
}

func maxProductBF(nums []int) int {
	maxValue := math.MinInt64
	dp := make([]int, len(nums))
	for start := 0; start < len(nums); start++ {
		for end := start; end < len(nums); end++ {
			if dp[end] != 0 {
				dp[end] = dp[end] / nums[start-1]
			} else {
				dp[end] = 1
				for i := start; i <= end; i++ {
					dp[end] *= nums[i]
				}
			}
			if dp[end] > maxValue {
				maxValue = dp[end]
			}
		}
	}
	return maxValue
}
