package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
	"AlgorithmWithGo/myLibs"
)

func ExampleDiameterOfBinaryTree() {
	fmt.Println(leetcode.DiameterOfBinaryTree(myLibs.MakeTreeNode("3,1,null,null,2")))
	fmt.Println(leetcode.DiameterOfBinaryTree(myLibs.MakeTreeNode("1,2,3")))
	fmt.Println(leetcode.DiameterOfBinaryTree(myLibs.MakeTreeNode("1,2,3,4,5")))
	fmt.Println(leetcode.DiameterOfBinaryTree(myLibs.MakeTreeNode("4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2")))
	// output:
	// 2
	// 2
	// 3
	// 8
}
