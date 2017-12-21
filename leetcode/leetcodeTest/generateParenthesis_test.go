package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleGenerateParenthesis() {
	fmt.Println(leetcode.GenerateParenthesis(3))
	// output:
	// [((())) (()()) (())() ()(()) ()()()]
}
