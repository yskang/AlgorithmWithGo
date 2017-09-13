package leetcode

func Rob1(nums []int) int {
	return rob1(nums)
}

func rob1(nums []int) int {
	robState := make([]int, len(nums))
	skipState := make([]int, len(nums))

	robState[0] = nums[0]
	skipState[0] = 0

	for i := 1 ; i < len(nums) ; i++ {
		robState[i] = skipState[i-1] + nums[i]
		skipState[i] = maxRob(robState[i-1], skipState[i-1])
	}

	return maxRob(robState[len(nums)-1], skipState[len(nums)-1])
}

func maxRob(a, b int) int {
	if a > b {
		return a
	}
	return b
}
