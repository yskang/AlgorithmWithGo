package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
	"AlgorithmWithGo/myLibs"
)

func ExampleMergeTree() {
	t1 := myLibs.MakeTreeNode("1,3,2,5")
	t2 := myLibs.MakeTreeNode("2,1,3,null,4,null,7")

	fmt.Println(leetcode.MergeTrees(t1, t2))

	// output:
	// 3,4,5,5,4,null,7
}