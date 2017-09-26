package leetcodeTests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleCombinationSumIV() {
	fmt.Println(leetcode.CombinationSum4([]int{1,2,3}, 4))
	fmt.Println(leetcode.CombinationSum4([]int{1,2,3}, 32))
	// output:
	// 7
	// 181997601
}