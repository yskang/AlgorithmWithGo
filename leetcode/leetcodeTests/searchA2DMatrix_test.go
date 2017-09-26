package leetcodeTests

import (
	"testing"
	"AlgorithmWithGo/leetcode"
)

func TestSearch2DMatrixSimple(t *testing.T) {
	input := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}

	target := 3

	expect := true

	result := leetcode.SearchMatrix(input, target)

	if result != expect {
		t.Error("Error, expect is ", expect, " but result is ", result)
	}
}