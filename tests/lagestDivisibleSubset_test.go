package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleLagestDivisibleSubset() {
	fmt.Println(leetcode.LargestDivisibleSubset([]int{2000000000}))
	fmt.Println(leetcode.LargestDivisibleSubset([]int{1,2,3}))
	fmt.Println(leetcode.LargestDivisibleSubset([]int{1,2,4,8}))
	fmt.Println(leetcode.LargestDivisibleSubset([]int{1,2,3,6,8,9,10,12,14}))
	// output:
	// [2000000000]
	// [1 2]
	// [1 2 4 8]
	// [1 2 6 12]
}
