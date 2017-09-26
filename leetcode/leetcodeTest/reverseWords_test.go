package leetcodeTest

import (
	"testing"
	"AlgorithmWithGo/leetcode"
)

func TestEmptyWord(t *testing.T) {
	result := leetcode.ReverseWords("")
	if result != "" {
		t.Error("Empty string error")
	}
}

func TestSingleWord(t *testing.T) {
	result := leetcode.ReverseWords("Let's")
	answer := "s'teL"

	if result != answer {
		t.Error("single string error")
	}
}

func TestSampleWord(t *testing.T) {
	result := leetcode.ReverseWords("Let's take LeetCode contest")
	answer := "s'teL ekat edoCteeL tsetnoc"

	if result != answer {
		t.Error("sample string error")
	}
}