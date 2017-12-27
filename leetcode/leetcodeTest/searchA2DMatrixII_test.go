package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleSearchA2DMatrixII() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(leetcode.SearchMatrixII(matrix, 5))
	fmt.Println(leetcode.SearchMatrixII(matrix, 20))
	fmt.Println(leetcode.SearchMatrixII([][]int{}, 0))
	fmt.Println(leetcode.SearchMatrixII([][]int{[]int{}}, 0))
	// output:
	// true
	// false
	// false
	// false
}
