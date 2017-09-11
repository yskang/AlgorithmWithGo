package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleMinimumPathSum() {
	fmt.Println(leetcode.MinPathSum([][]int{{0,1,2},{1,2,3}}))
	// output:
	// 6
}
