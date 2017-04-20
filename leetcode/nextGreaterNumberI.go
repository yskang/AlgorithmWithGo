package leetcode

func NextGreaterNumberI(findNums []int, nums []int) []int {
	return nextGreaterElementI(findNums, nums)
}

func nextGreaterElementI(findNums []int, nums []int) []int {
	bigNums := make(map[int]int)
	for start := 0 ; start < len(nums) ; start++ {
		for offset := 0 ; start + offset < len(nums) ; offset++ {
			if nums[start] < nums[start + offset] {
				bigNums[nums[start]] = nums[start + offset]
				break
			}
			if start + offset == len(nums) - 1 {
				bigNums[nums[start]] = -1
			}
		}
	}

	for i := 0 ; i < len(findNums) ; i++ {
		findNums[i] = bigNums[findNums[i]]
	}

	return findNums
}
