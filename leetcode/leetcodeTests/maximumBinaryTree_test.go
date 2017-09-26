package leetcodeTests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleMaximumBinaryTree() {
	fmt.Println(leetcode.ConstructMaximumBinaryTree([]int{3,2,1,6,0,5}))
	// output:
	// 6,3,5,null,2,0,null,null,1
}