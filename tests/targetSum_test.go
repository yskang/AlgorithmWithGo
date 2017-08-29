package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleTargetSum() {
	fmt.Println(leetcode.FindTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	// output:
	// 5
}