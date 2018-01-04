package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleSearch() {
	fmt.Println(leetcode.Search([]int{4, 5, 6, 7, 0, 1, 2}, 7))
	fmt.Println(leetcode.Search([]int{4, 5, 6, 7, 0, 1, 2}, 1))
	fmt.Println(leetcode.Search([]int{4, 5, 6, 7, 0, 1, 2}, 9))

	// output:
	// 3
	// 5
	// -1
}
