package leetcodeTests

import (
	"testing"
	"AlgorithmWithGo/leetcode"
)

func TestSampleWall(t *testing.T) {
	wall := [][]int{{1,2,2,1},
		{3,1,2},
		{1,3,2},
		{2,4},
		{3,1,2},
		{1,3,1,1}}
	result := leetcode.BrickWall(wall)
	if result != 2 {
		t.Error("sample wall error")
	}
}
