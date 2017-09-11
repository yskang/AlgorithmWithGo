package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleLongestIncreasingSubsequence() {
	fmt.Println(leetcode.LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	// output:
	// 4
}

