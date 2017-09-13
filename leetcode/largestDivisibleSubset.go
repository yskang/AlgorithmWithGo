package leetcode

import (
	"sort"
	"fmt"
)

func LargestDivisibleSubset(nums []int) []int {
	return largestDivisibleSubset(nums)
}

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	T := make([]int, len(nums))
	parent := make([]int, len(nums))

	m, ni := 0, 0

	for i := len(nums)-1 ; i >= 0 ; i-- {
		for j := i ; j < len(nums) ; j++ {
			if nums[j] % nums[i] == 0 && T[i] < 1 + T[j] {
				T[i] = 1 + T[j]
				parent[i] =j

				if T[i] > m {
					m = T[i]
					ni = i
				}
			}
		}
	}

	fmt.Println(T)
	fmt.Println(parent)

	ret := make([]int, 0)
	for i := 0 ; i < m ; i++ {
		ret = append(ret, nums[ni])
		ni = parent[ni]
	}
	return ret
}