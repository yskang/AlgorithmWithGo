package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleBulbSwitcherII() {
	fmt.Println(leetcode.FlipLights(1, 1))
	fmt.Println(leetcode.FlipLights(2, 1))
	fmt.Println(leetcode.FlipLights(3, 1))
	// output:
	// 2
	// 3
	// 4
}
