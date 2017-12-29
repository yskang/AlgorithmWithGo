package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"

	"github.com/yskang/leetcodeUtil/leetData"
)

func ExampleFlattenBinaryTreeToLinkedList() {
	tree := leetData.MakeTreeNode("1,2,5,3,4,null,6")
	leetcode.Flatten(tree)
	fmt.Println(tree)
	// output:
	// 1,null,2,null,3,null,4,null,5,null,6
}
