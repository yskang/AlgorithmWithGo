package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleSubsets() {
	fmt.Println(leetcode.Subsets([]int{1,2,3}))
	// output:
	// [[1] [2] [3] [1 2 3] [1 2] [1 3] [2 3] []]
}