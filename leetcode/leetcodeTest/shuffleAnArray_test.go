package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleShuffleAnArray() {
	nums := []int{1,2,3}
	solution := leetcode.ConstructorShuffle(nums)
	fmt.Println(solution.Shuffle())
	fmt.Println(solution.Reset())
	fmt.Println(solution.Shuffle())
	// output:
	//
}