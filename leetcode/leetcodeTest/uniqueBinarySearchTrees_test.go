package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleUniqueBinarySearchTrees() {
	fmt.Println(leetcode.NumTrees(4))
	fmt.Println(leetcode.NumTrees(5))
	fmt.Println(leetcode.NumTrees(6))

	// output:
	// 14
	// 42
	// 132
}