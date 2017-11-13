package leetcode

func FindMaximumXOR(nums []int) int {
	return findMaximumXOR(nums)
}

func findMaximumXOR(nums []int) int {
	max, mask := 0, 0

	for i := 31; i >= 0 ; i-- {
		mask = mask | 1 << uint(i)
		set := make(map[int]bool)
		for _, num := range nums {
			set[num & mask] = true
		}
		tmp := max | 1 << uint(i)
		for prefix := range set {
			if _, isExist := set[tmp ^ prefix]; isExist {
				max = tmp
				break
			}
		}
	}
	return max
}
