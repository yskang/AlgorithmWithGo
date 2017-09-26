package leetcodeTests

import (
	"testing"
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func TestNextGreaterNumberEmpty(t *testing.T) {
	input1 := []int{}
	input2 := []int{}
	expect := []int{}

	result := leetcode.NextGreaterNumberI(input1, input2)

	if len(result) != len(expect) {
		t.Error("error!")
	}

	for i := 0 ; i < len(result) ; i++ {
		if result[i] != expect[i] {
			t.Error("error!, expect is ", expect, "but result is ", result)
			return
		}
	}

	fmt.Println("success")
}

func TestSet1(t *testing.T) {
	input1 := []int{4, 1, 2}
	input2 := []int{1, 3, 4, 2}
	expect := []int{-1, 3, -1}

	result := leetcode.NextGreaterNumberI(input1, input2)

	if len(result) != len(expect) {
		t.Error("error!")
	}

	for i := 0 ; i < len(result) ; i++ {
		if result[i] != expect[i] {
			t.Error("error!, expect is ", expect, "but result is ", result)
			return
		}
	}

	fmt.Println("success")
}

func TestSet2(t *testing.T) {
	input1 := []int{2, 4}
	input2 := []int{1, 2, 3, 4}
	expect := []int{3, -1}

	result := leetcode.NextGreaterNumberI(input1, input2)

	if len(result) != len(expect) {
		t.Error("error!")
	}

	for i := 0 ; i < len(result) ; i++ {
		if result[i] != expect[i] {
			t.Error("error!, expect is ", expect, "but result is ", result)
			return
		}
	}

	fmt.Println("success")
}
