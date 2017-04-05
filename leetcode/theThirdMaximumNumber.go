package leetcode

import (
	"sort"
	"fmt"
)

func ThirdMax() {
	fmt.Println(thirdMax([]int{2,2,3,1}))
}

func thirdMax(nums []int) int {
	sort.Ints(nums)
	before := nums[len(nums) - 1]
	count := 1
	for i := len(nums) - 1 ; i >= 0 ; i-- {
		if before != nums[i] {
			count += 1
			before = nums[i]
			if count == 3 {
				return before
			}
		}
	}
	return nums[len(nums)-1]
}