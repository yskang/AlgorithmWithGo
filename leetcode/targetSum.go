package leetcode


func FindTargetSumWays(nums []int, S int) int {
	return findTargetSumWays(nums, S)
}


func findTargetSumWays(nums []int, S int) int {

	findCase(nums, S, 0, 0)

	return count
}
var count int
func findCase(nums []int, S int, i int, sum int) {
	if i == len(nums) {
		if sum == S {
			count += 1
		}
	} else {
		findCase(nums, S, i+1, sum + nums[i])
		findCase(nums, S, i+1, sum - nums[i])
	}
}