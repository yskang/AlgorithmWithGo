package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleMinCostClimbingStairs() {
	fmt.Println(leetcode.MinCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(leetcode.MinCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	// output:
	// 15
	// 6
}