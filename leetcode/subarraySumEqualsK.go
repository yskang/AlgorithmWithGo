package leetcode

// SubarraySum is 560th problem of leetcode
func SubarraySum(nums []int, k int) int {
	return subarraySum(nums, k)
}

func subarraySum(nums []int, k int) int {
	count := 0
	sums := make(map[int]int)
	sum := 0
	sums[0] = 1

	for _, num := range nums {
		sum += num
		if val, isExist := sums[sum-k]; isExist {
			count += val
		}
		sums[sum]++
	}

	return count
}
