package leetcodeTest_test

import (
	"AlgorithmWithGo/leetcode"
	"fmt"

	"github.com/yskang/leetcodeUtil/leetData"
)

func ExampleAddTwoNumbers() {
	fmt.Println(leetcode.AddTwoNumbers(leetData.InitListNode([]int{9, 8}), leetData.InitListNode([]int{1})))
	fmt.Println(leetcode.AddTwoNumbers(leetData.InitListNode([]int{5}), leetData.InitListNode([]int{5})))
	fmt.Println(leetcode.AddTwoNumbers(leetData.InitListNode([]int{2}), leetData.InitListNode([]int{2})))
	fmt.Println(leetcode.AddTwoNumbers(leetData.InitListNode([]int{2, 4, 3}), leetData.InitListNode([]int{5, 6, 4})))
	// output:
	// 0,9
	// 0,1
	// 4
	// 7,0,8
}
