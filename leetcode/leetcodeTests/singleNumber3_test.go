package leetcodeTests

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleSingleNumber3() {
	fmt.Println(leetcode.SingleNumber3([]int{1, 2, 1, 3, 2, 5}))
	// output:
	// [3 5]
}
