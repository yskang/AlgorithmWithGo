package leetcodeTests

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleShoppingOffers() {
	fmt.Println(leetcode.ShoppingOffers([]int{2,5}, [][]int{{3,0,5},{1,2,10}}, []int{3,2}))
	// output:
	// 14
}
