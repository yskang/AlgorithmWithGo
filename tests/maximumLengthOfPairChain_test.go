package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleMaximumLengthOfPairChain() {
	fmt.Println(leetcode.FindLongestChain([][]int{{1,2}, {2,3}, {3,4}}))
	// output:
	// 2
}
