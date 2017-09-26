package leetcodeTests

import (
	"fmt"
	"AlgorithmWithGo/myLibs"
	"AlgorithmWithGo/leetcode"
)

func ExampleAddOneRowToTree() {
	fmt.Println(leetcode.AddOneRowToTree(myLibs.MakeTreeNode("4,1,1,2,null,null,6,3,1,5"), 1, 2))
	fmt.Println(leetcode.AddOneRowToTree(myLibs.MakeTreeNode("4,2,6,3,1,5"), 1, 2))
	fmt.Println(leetcode.AddOneRowToTree(myLibs.MakeTreeNode("4,2,null,3,1"), 1, 3))
	fmt.Println(leetcode.AddOneRowToTree(myLibs.MakeTreeNode("1,2,3,4"), 5, 4))
	// output:
	// 4,1,1,1,null,null,1,2,null,null,6,3,1,5
	// 4,1,1,2,null,null,6,3,1,5
	// 4,2,null,1,1,3,null,null,1
	// 1,2,3,4,null,null,null,5,5
}
