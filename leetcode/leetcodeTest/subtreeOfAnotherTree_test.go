package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
	"AlgorithmWithGo/myLibs"
)

func ExampleIsSubtree() {
	fmt.Println(leetcode.IsSubtree(myLibs.MakeTreeNode("3,4,5,1,2"), myLibs.MakeTreeNode("4,1,2")))
	fmt.Println(leetcode.IsSubtree(myLibs.MakeTreeNode("3,4,5,1,2,null,null,null,null,0"), myLibs.MakeTreeNode("4,1,2")))
	// output:
	// true
	// false
}
