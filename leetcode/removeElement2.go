package main

import "fmt"

func main() {
	input := []int{1,2,3,4,5,4}
	size := removeElement(input, 4)
	fmt.Println(input, size)
}
func removeElement(nums []int, val int) int {
	count := 0

	for index := 0 ; index < len(nums) ;  {
		if nums[index] == val {
			for i := index ; i < len(nums) - count - 1 ; i++ {
				nums[i] = nums[i+1]
			}
			count += 1
			if index == len(nums) - count {
				fmt.Println(index)
				break
			}
		} else {
			index += 1
		}
	}

	return len(nums) - count
}
