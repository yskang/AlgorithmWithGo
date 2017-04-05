package leetcode

import (
	"fmt"
	"sort"
)

func MissingNumber() {
	fmt.Println(missingNumber([]int{0,1,2,3,4,6,7,8,9}))
}

func missingNumber(nums []int) int {
	sort.Ints(nums)

	for index, num := range(nums) {
		if index != num {
			return index
		}
	}

	return len(nums)
}
