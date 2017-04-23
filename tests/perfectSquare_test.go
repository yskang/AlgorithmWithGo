package tests

import (
	"testing"
	"AlgorithmWithGo/leetcode"
)

func TestPerfectSquareSimple9(t *testing.T) {
	input := 9
	expect := 1
	result := leetcode.PerfectSquare(input)

	if result != expect {
		t.Error("Error, expect is ", expect, " but result is ", result)
	}
	t.Log("test end")
}


func TestPerfectSquareSimple12(t *testing.T) {
	input := 12
	expect := 3
	result := leetcode.PerfectSquare(input)

	if result != expect {
		t.Error("Error, expect is ", expect, " but result is ", result)
	}
	t.Log("test end")
}

func TestPerfectSquareSimple13(t *testing.T) {
	input := 13
	expect := 2
	result := leetcode.PerfectSquare(input)

	if result != expect {
		t.Error("Error, expect is ", expect, " but result is ", result)
	}
}

func TestPerfectSquareSimple127(t *testing.T) {
	input := 127
	expect := 4
	result := leetcode.PerfectSquare(input)

	if result != expect {
		t.Error("Error, expect is ", expect, " but result is ", result)
	}
	t.Log("test end")
}