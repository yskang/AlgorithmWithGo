package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleDecodeString() {
	fmt.Println(leetcode.DecodeString("3[a]2[bc]"))
	fmt.Println(leetcode.DecodeString("3[a2[c]]"))
	fmt.Println(leetcode.DecodeString("2[abc]3[cd]ef"))
	// output:
	// aaabcbc
	// accaccacc
	// abcabccdcdcdef
}