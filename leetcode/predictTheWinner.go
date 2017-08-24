package leetcode

func PredictTheWinner(nums []int) bool {
	return totalScore_A_minus_B(nums, 0, len(nums)-1, 1) >= 0
}

func totalScore_A_minus_B(nums []int, start int, end int, turn int) int {
	if start == end {
		return turn * nums[start]
	}
	a := turn * nums[start] + totalScore_A_minus_B(nums, start + 1, end, -turn)
	b := turn * nums[end] + totalScore_A_minus_B(nums, start, end - 1, -turn)
	return turn * max(turn * a, turn * b)
}