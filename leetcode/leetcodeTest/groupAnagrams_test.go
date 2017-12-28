package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleGroupAnagrams() {
	fmt.Println(leetcode.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// output:
	// [[ate eat tea] [nat tan] [bat]]
}
