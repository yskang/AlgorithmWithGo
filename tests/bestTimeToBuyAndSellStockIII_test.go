package tests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleBestTimeToBuyAndSellStockIII() {
	fmt.Println(leetcode.MaxProfit3([]int{1,4,2}))
	fmt.Println(leetcode.MaxProfit3([]int{1,2}))
	fmt.Println(leetcode.MaxProfit3([]int{1,2,3,4,5}))
	fmt.Println(leetcode.MaxProfit3([]int{3,2,6,5,0,3}))
	// output:
	// 3
	// 1
	// 4
	// 7
}