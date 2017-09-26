package leetcodeTests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleBestTimeToBuyAndSellStockWithCooldown() {
	fmt.Println(leetcode.MaxProfit([]int{1,2,3,0,2}))
	// output:
	// 3
}
