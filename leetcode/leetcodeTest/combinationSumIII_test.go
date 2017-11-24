package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleCombinationSumIII()  {
	fmt.Println(leetcode.CombinationSum3(3, 7))
	fmt.Println(leetcode.CombinationSum3(3, 9))
	// output:
	// [[1 2 4]]
	// [[1 2 6] [1 3 5] [2 3 4]]
}