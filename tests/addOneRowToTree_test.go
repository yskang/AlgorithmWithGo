package tests

import (
	"fmt"
	"AlgorithmWithGo/myLibs"
	"AlgorithmWithGo/leetcode"
)

func ExampleAddOneRowToTree() {
	fmt.Println(myLibs.PrintTreeNode(leetcode.AddOneRowToTree(myLibs.MakeTreeNode("4,1,1,2,null,null,6,3,1,5"), 1, 2)))
	// output:
	// 4,1,1,1,null,null,1,2,null,null,6,3,1,5
}
