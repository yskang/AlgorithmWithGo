package main

import (
	"fmt"
)

func test() {
	fmt.Println(rob([]int{}))
	fmt.Println(rob([]int{1}))
	fmt.Println(rob([]int{1,2}))
	fmt.Println(rob([]int{1,2,3}))
	fmt.Println(rob([]int{1,2,3,4}))
	fmt.Println(rob([]int{1,2,3,4,5}))
	fmt.Println(rob([]int{10, 2, 3, 4, 3, 6, 7, 8, 3, 5, 10, 3, 4, 5, 7}))
	fmt.Println(rob([]int{114,117,207,117,235,82,90,67,143,146,53,108,200,91,80,223,58,170,110,236,81,90,222,160,165,195,187,199,114,235,197,187,69,129,64,214,228,78,188,67,205,94,205,169,241,202,144,240}))
}

func rob(nums []int) int {

	cache := make(map[int]int)

	return maxMoney(len(nums), nums, cache)
}

func maxMoney(numOfHouse int, nums []int, cache map[int]int) int {

	if val, exist := cache[numOfHouse]; exist {
		return val
	}

	if numOfHouse == 0 {
		return 0
	} else if numOfHouse == 1 {
		return nums[0]
	}

	a := maxMoney(numOfHouse-2, nums, cache) + nums[numOfHouse-1]
	b := maxMoney(numOfHouse-1, nums, cache)

	if a > b {
		cache[numOfHouse] = a
		return a
	}
	cache[numOfHouse] = b
	return b
}