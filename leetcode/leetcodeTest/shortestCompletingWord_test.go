package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleShortestCompletingWord() {
	fmt.Println(leetcode.ShortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
	fmt.Println(leetcode.ShortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"}))
	// output:
	// steps
	// pest
}