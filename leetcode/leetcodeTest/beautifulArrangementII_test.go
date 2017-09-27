package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleBeautifulArrangementII() {
	fmt.Println(leetcode.ConstructArray(3, 1))
	fmt.Println(leetcode.ConstructArray(3, 2))
	fmt.Println(leetcode.ConstructArray(5, 2))
	fmt.Println(leetcode.ConstructArray(10, 9))
	fmt.Println(leetcode.ConstructArray(10, 4))
	fmt.Println(leetcode.ConstructArray(10, 3))

	// output:
	// [1 2 3]
	// [1 3 2]
	// [1 3 2 4 5]
	// [1 10 2 9 3 8 4 7 5 6]
}
