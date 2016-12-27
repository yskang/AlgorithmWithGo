package main

import (
	"fmt"
)

func main() {
	input := []int{2, 1, 1, 2, 1, 1}
	size := removeElement(input, 1)
	fmt.Println(input, size)
}

func removeElement(nums []int, val int) int {

	replacer := -1
	searchEnd := len(nums)
	backwardStart := len(nums) - 1
	valid := false

	for i := 0 ; i < searchEnd ; i++ {
		if nums[i] == val {
			replacer, backwardStart, valid = getReplacer(nums, val, backwardStart, i)

			if valid {
				nums[i] = replacer
				searchEnd = backwardStart + 1
			} else {
				searchEnd = i
			}

		}
	}

	return searchEnd
}

func getReplacer(nums []int, removeValue int, start int, end int) (int, int, bool) {
	for i := start ; i > end ; i-- {
		if nums[i] != removeValue {
			return nums[i], i - 1, true
		}
	}

	return 0, end, false
}
