package tests

import (
	"testing"
	"AlgorithmWithGo/leetcode"
)

func TestRotateImageSimple(t *testing.T) {
	input := [][]int{{1,2,3},
		{4,5,6},
		{7,8,9},}
	expect := [][]int{{7,4,1},
		{8,5,2},
		{9,6,3},}
	result := leetcode.RotateImage(input)

	for i := 0 ; i < len(input) ; i++ {
		for j := 0 ; j < len(input[0]) ; j++ {
			if result[i][j] != expect[i][j] {
				t.Error("error!, expectd: ", expect, ", but result is ", result, ".")
			}
		}
	}
}