package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleFindDuplicateNum() {
	fmt.Println(leetcode.FindDuplicate([]int{1,3,4,2,2}))
	// output:
	// 2
}