package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleOnesAndZeros() {
	fmt.Println(leetcode.FindMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	// output:
	// 4
}
