package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/myLibs"
	"AlgorithmWithGo/leetcode"
)

func ExamplePrintBinaryTree() {
	fmt.Println(leetcode.PrintTree(myLibs.MakeTreeNode("1,2")))
	fmt.Println(leetcode.PrintTree(myLibs.MakeTreeNode("1,2,5,3,null,null,6,7,null,null,8")))
	// output:
	// [[ 1 ] [2  ]]
	// [[       1       ] [   2        5   ] [ 3            6 ] [7              8]]

}
