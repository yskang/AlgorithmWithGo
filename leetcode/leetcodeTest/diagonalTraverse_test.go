package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleFindDiagonalOrder() {
	fmt.Println(leetcode.FindDiagonalOrder([][]int{
		{1, 2, 3},
	}))
	// output:
	// [1 2 3]
}
