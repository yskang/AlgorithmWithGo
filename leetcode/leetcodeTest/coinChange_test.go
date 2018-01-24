package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleCoinChange() {
	fmt.Println(leetcode.CoinChange([]int{227, 99, 328, 299, 42, 322}, 9847))
	fmt.Println(leetcode.CoinChange([]int{1, 2147483647}, 2))
	fmt.Println(leetcode.CoinChange([]int{2}, 1))
	fmt.Println(leetcode.CoinChange([]int{1}, 0))
	fmt.Println(leetcode.CoinChange([]int{1, 2, 5}, 11))
	// output:
	// 2
	// -1
	// 0
	// 3
}
