package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleTaskScheduler_test() {
	//fmt.Println(leetcode.LeastInterval([]byte{'A','A','A','B','B','B'}, 2))
	fmt.Println(leetcode.LeastInterval([]byte{'A','A','A','A','A','A','B','C','D','E','F','G'}, 2))
	// output:
	// 8
}
