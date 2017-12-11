package leetcode

func Permute(nums []int) [][]int {
	return permute(nums)
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	getPermute(nums, 0, len(nums), len(nums), &ans)
	return ans
}

func getPermute(nums []int, depth int, n int, k int, ans *[][]int) {
	if depth == k {
		*ans = append(*ans, append([]int{}, nums...))
		return
	}

	for i := depth ; i < n ; i++ {
		nums[i], nums[depth] = nums[depth], nums[i]
		getPermute(nums, depth+1, n, k, ans)
		nums[i], nums[depth] = nums[depth], nums[i]
	}
}