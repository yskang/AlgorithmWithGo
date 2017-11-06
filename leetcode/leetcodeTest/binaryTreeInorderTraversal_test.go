package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
	"AlgorithmWithGo/myLibs"
)

func ExampleInorderTraversal() {
	fmt.Println(leetcode.InorderTraversal(myLibs.MakeTreeNode("1,null,2,3")))
	// output:
	// [1 3 2]
}