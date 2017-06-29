package leetcode

import (
	"strconv"
	"math"
)

func OptimalDivision(nums []int) string {
	return  optimalDivision(nums)
}

func optimalDivision(nums []int) string {
	t := optimal(nums, 0, len(nums)-1, "")
	return t.maxStr
}

type T struct {
	maxVal float64
	minVal float64
	minStr string
	maxStr string
}

func optimal(nums []int, start int, end int, res string) T {
	t := T{}
	if start == end {
		t.maxVal = float64(nums[start])
		t.minVal = float64(nums[start])
		t.minStr = strconv.Itoa(nums[start])
		t.maxStr = strconv.Itoa(nums[start])
		return t
	}
	t.minVal = math.MaxFloat64
	t.maxVal = math.SmallestNonzeroFloat64
	t.minStr, t.maxStr = "", ""
	for i := start ; i < end ; i++ {
		left := optimal(nums, start, i, "")
		right := optimal(nums, i + 1, end, "")
		if t.minVal > left.minVal / right.maxVal {
			t.minVal = left.minVal / right.maxVal
			if i + 1 != end {
				t.minStr = left.minStr + "/" + "(" + right.maxStr + ")"
			} else {
				t.minStr = left.minStr + "/" + right.maxStr
			}
		}
		if t.maxVal < left.maxVal / right.minVal {
			t.maxVal = left.maxVal /right.minVal
			if i + 1 != end {
				t.maxStr = left.maxStr + "/" + "(" + right.minStr + ")"
			} else {
				t.maxStr = left.maxStr + "/" + right.minStr
			}
		}
	}
	return t
}