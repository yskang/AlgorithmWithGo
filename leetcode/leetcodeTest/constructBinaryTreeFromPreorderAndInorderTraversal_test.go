package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleBuildTree() {
	fmt.Println(leetcode.BuildTree([]int{7, 10, 4, 3, 1, 2, 8, 11}, []int{4, 10, 3, 1, 7, 11, 8, 2}))
	// output:
	// 7,10,2,4,3,8,null,null,null,null,1,11
}
