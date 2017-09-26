package leetcodeTests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleSingleElementInASortedArray() {
	fmt.Println(leetcode.SingleElementInASortedArray([]int{1,1,2,3,3,4,4,8,8}))
	fmt.Println(leetcode.SingleElementInASortedArray([]int{1,1,2}))
	// Output:
	// 2
	// 2
}
