package leetcode

func FindDuplicate(nums []int) int {
	return findDuplicateNum(nums)
}

func findDuplicateNum(nums []int) int {
	tortoise, hare := nums[0], nums[nums[0]]

	for tortoise != hare{
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
	}

	ptr1, ptr2 := 0, tortoise
	for ptr1 != ptr2 {
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}
	return ptr1
}
