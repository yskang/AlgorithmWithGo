package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleTeemoAttacking() {
	fmt.Println(leetcode.TeemoAttacking([]int{1,4}, 2))
	fmt.Println(leetcode.TeemoAttacking([]int{1,2}, 2))
	fmt.Println(leetcode.TeemoAttacking([]int{1,2,3,4,5}, 5))
	// Output:
	// 4
	// 3
	// 9
}
