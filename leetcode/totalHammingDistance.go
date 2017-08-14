package leetcode

func TotalHammingDistance(nums []int) int {
	return totalHammingDistance(nums)
}

func totalHammingDistance(nums []int) int {
	sum := 0
	shiftCount := uint(0)

	for shiftCount < 32 {
		zeros := 0
		for _, num := range nums {
			num = num >> shiftCount
			if num & 0x01 == 0 {
				zeros += 1
			}
		}
		sum += zeros * (len(nums) - zeros)
		shiftCount += 1
	}

	return sum
}