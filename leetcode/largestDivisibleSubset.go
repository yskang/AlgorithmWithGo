package leetcode

import (
	"sort"
)

func LargestDivisibleSubset(nums []int) []int {
	return largestDivisibleSubset(nums)
}

func largestDivisibleSubset(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	sort.Ints(nums)

	seeds := make([][]int, 1)
	seeds[0] = append([]int{}, nums[0])

	for i := 1 ; i < len(nums) ; i++ {
		isChildOfOtherSeed := false
		for seedIndex, children := range seeds {
			for j := len(children)-1 ; j >= 0 ; j-- {
				if nums[i] % children[j] == 0 {
					if j == len(children)-1 {
						seeds[seedIndex] = append(seeds[seedIndex], nums[i])
					} else {
						temp := append([]int{}, seeds[seedIndex][:j+1]...)
						temp = append(temp, nums[i])
						seeds = append(seeds, []int{})
						seeds[len(seeds)-1] = append(seeds[len(seeds)-1], temp...)
					}
					isChildOfOtherSeed = true
				}
				if isChildOfOtherSeed {
					break
				}
			}
		}
		if !isChildOfOtherSeed {
			seeds = append(seeds, []int{nums[i]})
		}
	}

	maxL := -1
	maxSeed := 0
	for seed, children := range seeds {
		if maxL < len(children) {
			maxL = len(children)
			maxSeed = seed
		}
	}

	return seeds[maxSeed]
}