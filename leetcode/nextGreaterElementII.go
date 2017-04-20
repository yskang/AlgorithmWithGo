package leetcode

func NextGreaterElementII(nums []int) []int {
	return nextGreaterElements(nums)
}

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	newNums := make([]int, 0)
	newNums = append(newNums, nums...)
	newNums = append(newNums, nums...)

	for start := 0 ; start < len(nums) ; start++ {
		for offset := 0 ; offset < len(nums) ; offset++ {
			if newNums[start + offset] > newNums[start] {
				newNums[start] = newNums[start + offset]
				break
			}
			if offset == len(nums) - 1 {
				newNums[start] = -1
			}
		}
	}

	return newNums[0:len(nums)]
}
