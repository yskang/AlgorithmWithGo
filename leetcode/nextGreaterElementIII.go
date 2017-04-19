package leetcode

import (
	"strconv"
	"sort"
	"math"
)

func NextGreaterElementIII(n int) int {
	return nextGreaterElement(n)
}

func nextGreaterElement(n int) int {
	numStr := strconv.Itoa(n)

	prevDigit := numStr[len(numStr) - 1]
	for i := len(numStr) - 2 ; i >= 0 ; i-- {
		if numStr[i] < prevDigit {
			return greaterDigit(numStr, i)
		}
		prevDigit = numStr[i]
	}

	return -1
}

func greaterDigit(numStr string, index int) int {
	minDiff := 100
	nextNumIndex := 0
	for i := index+1 ; i < len(numStr) ; i++ {
		diff := int(numStr[i] - numStr[index])
		if diff < minDiff && diff > 0 {
			minDiff = diff
			nextNumIndex = i
		}
	}

	nums := make([]int, 0)
	baseNumber, _ := strconv.Atoi(string(numStr[index]))
	nums = append(nums, baseNumber)
	for i := index + 1 ; i < len(numStr) ; i++ {
		if i != nextNumIndex {
			num, _ := strconv.Atoi(string(numStr[i]))
			nums = append(nums, num)
		}
	}
	sort.Ints(nums)


	resultNumStr := numStr[:index] + string(numStr[nextNumIndex])

	for _, number := range nums {
		resultNumStr += strconv.Itoa(number)
	}

	resultNum, _ := strconv.Atoi(resultNumStr)

	if resultNum > math.MaxInt32 {
		return -1
	}

	return resultNum
}