package leetcodeTest_test

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleWordSearch() {
	fmt.Println(leetcode.Exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}, "ABCCED"))
	// output:
	// true
}
