package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleKeysKeyboard() {
	fmt.Println(leetcode.MinSteps(3))
	fmt.Println(leetcode.MinSteps(50))
	fmt.Println(leetcode.MinSteps(125))
	fmt.Println(leetcode.MinSteps(10000000))
	// output:
	// 3
	// 12
	// 15
	// 49
}
