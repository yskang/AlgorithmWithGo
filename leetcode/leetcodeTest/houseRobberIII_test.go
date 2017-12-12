package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
	"AlgorithmWithGo/myLibs"
)

func ExampleRobIII() {
	fmt.Println(leetcode.RobIII(myLibs.MakeTreeNode("3,2,3,null,3,null,1")))
	fmt.Println(leetcode.RobIII(myLibs.MakeTreeNode("3,4,5,1,3,null,1")))
	// output:
	// 7
	// 9
}