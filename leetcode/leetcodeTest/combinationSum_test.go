package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleCombinationSum() {
	fmt.Println(leetcode.CombinationSum([]int{2, 3, 4, 7}, 7))
	// output:
	// [[2 2 3] [3 4] [7]]
}
