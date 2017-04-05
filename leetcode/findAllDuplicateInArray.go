package leetcode

import (
	"math"
)

func FindAllDuplicatInArray(nums []int) []int {
	return findDuplicates(nums)
}

func findDuplicates(nums []int) []int {
	result := make([]int, 0)

	for _, num := range(nums) {
		if nums[int(math.Abs(float64(num))-1)] < 0 {
			result = append(result, int(math.Abs(float64(num))))
		} else {
			nums[int(math.Abs(float64(num))-1)] *= -1
		}
	}

	return result
}
