package leetcode

func ProductOfArrayExceptSelf(nums []int) []int {
	return productExceptSelf(nums)
}

func productExceptSelf(nums []int) []int {
	productOfAll := 1
	zeroCount := 0

	for _, num := range nums {
		if num != 0 {
			productOfAll *= num
		} else {
			zeroCount += 1
		}
	}

	output := make([]int, len(nums))

	if zeroCount == 1 {
		for i, num := range nums {
			if num == 0 {
				output[i] = productOfAll
			} else {
				output[i] = 0
			}
		}
	} else if zeroCount >= 2 {
	} else {
		for i, num := range nums {
			output[i] = productOfAll / num
		}
	}

	return output
}