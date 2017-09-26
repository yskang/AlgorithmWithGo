package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleTopKFrequentElement() {
	fmt.Println(leetcode.TopKFrequentElements([]int{1,1,1,2,2,3}, 2))

	// output:
	// [1 2]
}
