package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleCourseSchedule() {
	fmt.Println(leetcode.CanFinish(2, [][]int{{1, 0}}))
	fmt.Println(leetcode.CanFinish(2, [][]int{{1, 0}, {0, 1}}))
	// output:
	// true
	// false
}
