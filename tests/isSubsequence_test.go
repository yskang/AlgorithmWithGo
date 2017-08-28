package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleIsSubsequence() {
	fmt.Println(leetcode.IsSubsequence("abbc", "ahbgdc"))
	fmt.Println(leetcode.IsSubsequence("axc", "ahbgdc"))

	// output:
	// false
	// false
}