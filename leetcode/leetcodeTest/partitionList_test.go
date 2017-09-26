package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
	"AlgorithmWithGo/myLibs"
)

func ExamplePartitionList() {
	fmt.Println(leetcode.Partition(myLibs.InitListNode([]int{1,4,3,2,5,2}), 3))
	// output:
	// 1,2,2,4,3,5
}
