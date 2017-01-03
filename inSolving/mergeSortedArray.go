package main

import (
	"fmt"
)

func main() {
	merge([]int{1,2,3,4,5,9}, 5, []int{2,5,7,8}, 4)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	i := 0

	for {
		fmt.Println(nums1, nums2)
		fmt.Println(i, "th", j, "th")
		fmt.Println(nums1[i], nums2[j])
		if nums1[i] < nums2[j] {
			fmt.Println("nums2 is big")
			i += 1
		} else {
			fmt.Println("nums1 is big")
			fmt.Println(nums1[:i+2], "+", nums1[i+1:])
			nums1 = append(nums1[:i+2], nums1[i+1:]...)

			nums1[i+1] = nums2[j]
			j += 1
			i += 1
		}


		if i >= len(nums1) {
			nums1 = append(nums1, nums2[j:]...)
			break
		} else if j == n {
			break
		}
	}


	fmt.Println(nums1)
}

