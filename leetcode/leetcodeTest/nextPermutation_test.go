package leetcodeTest_test

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleNextPurmutation() {
	fmt.Println(leetcode.NextPermutation([]int{1, 2, 3}))
	fmt.Println(leetcode.NextPermutation([]int{3, 2, 1}))
	fmt.Println(leetcode.NextPermutation([]int{1, 1, 5}))
	// output:
	// [1 3 2]
	// [1 2 3]
	// [1 5 1]
}
