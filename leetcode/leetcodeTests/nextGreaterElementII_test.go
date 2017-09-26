package leetcodeTests

import (
	"testing"
	"AlgorithmWithGo/leetcode"
)

func TestEmptyArray(t *testing.T) {
	result := leetcode.NextGreaterElementII([]int{})

	if len(result) != 0 {
		t.Error("empty array error")
	}
}

func TestSimpleExample(t *testing.T) {
	input := []int{1, 2, 1}
	expect := []int{2, -1, 2}
	result := leetcode.NextGreaterElementII(input)

	if len(result) != 3 {
		t.Error("Error result size")
	}

	for i := 0 ; i < len(result) ; i++ {
		if result[i] != expect[i] {
			t.Error("Error! expect is ", expect, "but result is ", result)
			return
		}
	}
}

func TestSimpleExample2(t *testing.T) {
	input := []int{5,4,3,2,1}
	expect := []int{-1,5,5,5,5}
	result := leetcode.NextGreaterElementII(input)

	if len(result) != 5 {
		t.Error("Error result size")
	}

	for i := 0 ; i < len(result) ; i++ {
		if result[i] != expect[i] {
			t.Error("Error! expect is ", expect, "but result is ", result)
			return
		}
	}
}
