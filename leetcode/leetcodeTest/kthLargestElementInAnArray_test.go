package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleFindKthLargest() {
	fmt.Println(leetcode.FindKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	// output:
	// 5
}
