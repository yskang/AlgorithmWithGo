package leetcodeTests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleTriangle() {
	fmt.Println(leetcode.MinimumTotal([][]int{{2}, {3,4}, {6,5,7}, {4,1,8,3}}))
	// output:
	// 11
}
