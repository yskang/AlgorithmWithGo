package leetcode

func TotalHammingDistance(nums []int) int {
	return totalHammingDistance(nums)
}

func totalHammingDistance(nums []int) int {
	sum := uint32(0)

	for index := 0 ; index < len(nums) ; index++ {
		for i := index + 1 ; i < len(nums) ; i++ {
			sum += getHammingDistance(nums[index], nums[i])
		}
	}

	return int(sum)
}

func getHammingDistance(f int, s int) uint32 {
	a := uint32(f)
	b := uint32(s)
	xor := a ^ b
	sum := uint32(0)

	for i := 0 ; i < 32 ; i++ {
		sum += (xor >> uint32(i)) & 1
	}

	return sum
}
