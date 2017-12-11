package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleFindUnsortedSubarray() {
	fmt.Println(leetcode.FindUnsortedSubarray([]int{2,6,4,8,10,9,15}))
	fmt.Println(leetcode.FindUnsortedSubarray([]int{1,2,3,4}))
	fmt.Println(leetcode.FindUnsortedSubarray([]int{1,3,2,2,2}))
	fmt.Println(leetcode.FindUnsortedSubarray([]int{1,2,3,3,3}))
	fmt.Println(leetcode.FindUnsortedSubarray([]int{2,3,3,2,4}))
	fmt.Println(leetcode.FindUnsortedSubarray([]int{1,2,4,5,3}))
	// output:
	// 5
	// 0
	// 4
	// 0
	// 3
	// 3
}
