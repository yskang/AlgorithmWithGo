package main

import "fmt"

func main() {
	x := []int{1,1,2,2,2,2}
	l := removeDuplicates(x)
	fmt.Println(x)
	fmt.Println(x[:l])
}

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	count := 0

	for i := 0 ; i + count < len(nums) - 1; {
		if nums[i] == nums[i+1] {
			fmt.Println("remove", i, "th", nums[i])
			for j := i ; j < len(nums) - count - 1 ; j++ {
				nums[j] = nums[j+1]
			}
			count += 1
		} else {
			i += 1
		}
	}

	return len(nums) - count
}
