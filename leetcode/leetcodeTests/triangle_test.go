package leetcodeTests

import (
	"testing"
	"AlgorithmWithGo/leetcode"
)

func TestTriangle(t *testing.T) {
	result := leetcode.Triangle(9)
	const expect = 3

	if result != 3 {
		t.Error("Error, expected is ", expect, " but result is ", result)
	}
}
