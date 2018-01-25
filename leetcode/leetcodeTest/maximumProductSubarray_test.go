package leetcodeTest_test

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleMaximumPruductSubarray() {
	fmt.Println(leetcode.MaxProduct([]int{-2, 0, 2, 4, 0, 1}))
	fmt.Println(leetcode.MaxProduct([]int{3, -1, 4}))
	fmt.Println(leetcode.MaxProduct([]int{2, 3, -2, 4}))
	// output:
	// 4
	// 6
}
