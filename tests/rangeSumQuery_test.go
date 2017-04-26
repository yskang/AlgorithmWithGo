package tests

import (
	"testing"
	"AlgorithmWithGo/leetcode"
)

func TestRangeSumQuerySimple(t *testing.T) {
	input := []int{1, 3, 5}

	obj := leetcode.NumArrayConstructor(input);

	expect1 := 9
	result1 := obj.SumRangeII(0, 2)

	obj.Update(1, 2);
	expect2 := 8
	result2 := obj.SumRangeII(0, 2);


	if result1 != expect1 {
		t.Error("Error! expect is ", expect1, " but, result is ", result1)
	}

	if result2 != expect2 {
		t.Error("Error! expect is ", expect2, " but, result is ", result2)
	}
}
