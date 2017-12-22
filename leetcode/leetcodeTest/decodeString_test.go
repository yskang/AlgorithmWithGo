package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleDecodeString() {
	fmt.Println(leetcode.DecodeString("3[a]2[bc]"))
	fmt.Println(leetcode.DecodeString("3[a2[c]]"))
	fmt.Println(leetcode.DecodeString("2[abc]3[cd]ef"))
	fmt.Println(leetcode.DecodeString("2[2[b]]"))
	// output:
	// aaabcbc
	// accaccacc
	// abcabccdcdcdef
	// bbbb
}
