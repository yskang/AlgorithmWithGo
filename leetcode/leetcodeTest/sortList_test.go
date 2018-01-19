package leetcodeTest_test

import (
	"AlgorithmWithGo/leetcode"
	"fmt"

	"github.com/yskang/leetcodeUtil/leetData"
)

func ExampleSortList() {
	fmt.Println(leetcode.SortList(leetData.InitListNode([]int{4, 10, 6, 8, 7, 9, 1, 5, 3, 2})))
	// output:
	// 1,2,3,4,5,6,7,8,9,10
}
