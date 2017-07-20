package leetcode

func SingleNumber3(nums []int) []int {
	return singleNumber(nums)
}

func singleNumber(nums []int) []int {
	twoNums := 0
	for _, num := range nums {
		twoNums ^= num
	}

	bitIndex := 0
	for {
		if twoNums & 0x01 == 1 {
			break
		}
		twoNums >>= 1
		bitIndex += 1
	}

	firstNum := 0
	secondNum := 0
	for _, num := range nums {
		if (num >> uint(bitIndex)) & 0x01 == 1 {
			firstNum ^= num
		} else {
			secondNum ^= num
		}
	}

	return []int{firstNum, secondNum}
}
