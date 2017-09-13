package leetcode

func Rob2(nums []int) int {
return rob2(nums)
}

func rob2(nums []int) int {
	if len(nums) < 1 {
		return 0
	} else if len(nums) < 2 {
		return nums[0]
	}

	robState := make([]int, len(nums))
	skipState := make([]int, len(nums))

	// rob first house case
	robState[0] = nums[0]
	skipState[0] = 0

	for i := 1 ; i < len(nums) ; i++ {
		robState[i] = skipState[i-1] + nums[i]
		skipState[i] = maxRob2(robState[i-1], skipState[i-1])
	}
	c1 := skipState[len(nums)-1]

	// skip first house case
	robState[0] = 0
	skipState[0] = 0

	for i := 1 ; i < len(nums) ; i++ {
		robState[i] = skipState[i-1] + nums[i]
		skipState[i] = maxRob2(robState[i-1], skipState[i-1])
	}
	c2 := maxRob2(robState[len(nums)-1], skipState[len(nums)-1])

	return maxRob2(c1, c2)
}

func maxRob2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
