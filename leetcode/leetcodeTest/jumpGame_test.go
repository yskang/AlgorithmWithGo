package leetcodeTest_test

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleJumpGame() {
	fmt.Println(leetcode.CanJump([]int{0}))
	fmt.Println(leetcode.CanJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(leetcode.CanJump([]int{3, 2, 1, 0, 4}))
	// output:
	// true
	// true
	// false
}
