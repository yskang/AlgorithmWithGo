package tests

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleOptimalDivision() {
	fmt.Println(leetcode.OptimalDivision([]int{1000,100,10,2}))
	// Output:
	// 1000/(100/10/2)
}
