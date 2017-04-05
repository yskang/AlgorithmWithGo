package leetcode

import "fmt"

func ContainNearbyDuplicate() {
	fmt.Println(containsNearbyDuplicate([]int{99, 99}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 || nums == nil || len(nums) ==  1 {
		return false
	}

	if len(nums) <= k {
		k = len(nums) - 1
	}

	dict := make(map[int]int)

	for i := 0 ; i < k + 1 ; i++ {
		if _, exist := dict[nums[i]]; exist{
			return true
		}
		dict[nums[i]] += 1
	}

	fmt.Println(dict)

	for i := 1 ; i < len(nums) - k ; i++ {
		dict[nums[i-1]] -= 1
		if value, exist := dict[nums[i + k]]; exist && value == 1{
			fmt.Println(i, k, nums[i + k] )
			return true
		}
		dict[nums[i+k]] += 1
		fmt.Println(dict)
	}

	return false
}
