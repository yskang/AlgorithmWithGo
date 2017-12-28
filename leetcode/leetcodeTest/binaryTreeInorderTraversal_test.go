package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"

	"github.com/yskang/leetcodeUtil/leetData"
)

func ExampleInorderTraversal() {
	fmt.Println(leetcode.InorderTraversal(leetData.MakeTreeNode("1,null,2,3")))
	// output:
	// [1 3 2]
}
