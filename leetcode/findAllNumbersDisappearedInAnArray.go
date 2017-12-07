package leetcode

func FindDisappearedNumbers(nums []int) []int {
	return findDisappearedNumbers(nums)
}

func findDisappearedNumbers(nums []int) []int {
	ans := make([]int, 0)
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = nums[num-1] * -1
		}
	}

	for index, num := range nums {
		if num > 0 {
			ans = append(ans, index+1)
		}
	}

	return ans
}