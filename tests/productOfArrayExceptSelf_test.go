package tests

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleProductOfArrayExceptSelf() {
	fmt.Println(leetcode.ProductOfArrayExceptSelf([]int{1,2,3,4}))
	fmt.Println(leetcode.ProductOfArrayExceptSelf([]int{0,0}))
	// output:
	// [24 12 8 6]
	// [0 0]
}
