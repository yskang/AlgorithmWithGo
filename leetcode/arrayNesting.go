package leetcode

func ArrayNesting(nums []int) int {
	return arrayNesting(nums)
}

func arrayNesting(nums []int) int {
	lengthCheck := make([]int, len(nums))
	maxLen := 0
	length := 0

	for _, num := range nums {
		if lengthCheck[num] == 0 {
			for {
				length += 1
				lengthCheck[num] = 1

				num = nums[num]
					if lengthCheck[num] == 1 {
						break
				}
			}

			if maxLen < length {
				maxLen = length
			}

			length = 0
		}
	}

	return maxLen
}
