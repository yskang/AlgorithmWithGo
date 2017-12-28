package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"

	"github.com/yskang/leetcodeUtil/treeNode"
)

func ExampleInorderTraversal() {
	fmt.Println(leetcode.InorderTraversal(treeNode.MakeTreeNode("1,null,2,3")))
	// output:
	// [1 3 2]
}
