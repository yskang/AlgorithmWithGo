package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleSearchForARange() {
	fmt.Println(leetcode.SearchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	// output:
	// [3 4]
}
