package leetcode

func WiggleMaxLength(nums []int) int {
	return wiggleMaxLength(nums)
}

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	zipNums := make([]int, 1)
	zipNums[0] = nums[0]
	for _, num := range nums[1:] {
		if zipNums[len(zipNums)-1] != num {
			zipNums = append(zipNums, num)
		}
	}

	if len(zipNums) < 2 {
		return len(zipNums)
	} else if len(zipNums) == 2 && zipNums[0] != zipNums[1] {
		return 2
	}

	diffSeq := make([]bool, 1)
	diffSeq[0] = zipNums[1] - zipNums[0] > 0
	prev := zipNums[1]
	for i := 2 ; i < len(zipNums) ; i++ {
		if zipNums[i] - prev < 0 {
			if diffSeq[len(diffSeq)-1] != false {
				diffSeq = append(diffSeq, false)
			}
		} else {
			if diffSeq[len(diffSeq)-1] != true {
				diffSeq = append(diffSeq, true)
			}
		}
		prev = zipNums[i]
	}

	return len(diffSeq)+1
}