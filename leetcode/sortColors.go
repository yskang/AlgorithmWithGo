package leetcode

import (
	"log"
)

//SortColors sort the colors
func SortColors(nums []int) []int {
	sortColors(nums)
	return nums
}

// counting method
func sortColors_(nums []int) {
	r, w, b := 0, 0, 0
	for _, num := range nums {
		switch num {
		case 0:
			r++
		case 1:
			w++
		case 2:
			b++
		}
	}
	for i := 0; i < r; i++ {
		nums[i] = 0
	}
	for i := r; i < r+w; i++ {
		nums[i] = 1
	}
	for i := r + w; i < r+w+b; i++ {
		nums[i] = 2
	}
}

// swap method, using constant memory
func sortColors(nums []int) {
	log.Println(nums)
	zero, two, i := 0, len(nums)-1, 0
	for {
		if nums[i] == 0 {
			nums[zero], nums[i] = nums[i], nums[zero]
			log.Println(nums, "swap", zero, i, "now zero is", zero+1)
			zero++
			i++
		} else if nums[i] == 2 {
			if i >= two {
				break
			}
			nums[two], nums[i] = nums[i], nums[two]
			log.Println(nums, "swap", i, two, "now tow is", two-1)
			two--
		} else {
			i++
		}
		if zero == len(nums) || two == -1 || i == len(nums) {
			break
		}
	}
}
