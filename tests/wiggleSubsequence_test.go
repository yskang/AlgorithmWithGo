package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleWiggleSubsequence() {
	fmt.Println(leetcode.WiggleMaxLength([]int{0,0}))
	fmt.Println(leetcode.WiggleMaxLength([]int{1,7,4,9,2,5}))
	fmt.Println(leetcode.WiggleMaxLength([]int{1,17,5,10,13,15,10,5,16,8}))
	fmt.Println(leetcode.WiggleMaxLength([]int{1,2,3,4,5,6,7,8,9}))
	// output:
	// 1
	// 6
	// 7
	// 2
}
