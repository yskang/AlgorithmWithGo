package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleSortColors() {
	fmt.Println(leetcode.SortColors([]int{0, 1, 2, 0, 1, 2}))
	fmt.Println(leetcode.SortColors([]int{0}))
	fmt.Println(leetcode.SortColors([]int{1}))
	// output:
	// [0 0 1 1 2 2]
	// [0]
	// [1]
}
