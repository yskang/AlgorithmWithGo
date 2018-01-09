package leetcodeTest_test

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
	"strings"
)

func ExampleWordBreak() {
	fmt.Println(leetcode.WordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(leetcode.WordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
	fmt.Println(leetcode.WordBreak("cars", []string{"car", "ca", "rs"}))
	fmt.Println(leetcode.WordBreak("bccdbacdbdacddabbaaaadababadad", []string{"cbc", "bcda", "adb", "ddca", "bad", "bbb", "dad", "dac", "ba", "aa", "bd", "abab", "bb", "dbda", "cb", "caccc", "d", "dd", "aadb", "cc", "b", "bcc", "bcd", "cd", "cbca", "bbd", "ddd", "dabb", "ab", "acd", "a", "bbcc", "cdcbd", "cada", "dbca", "ac", "abacd", "cba", "cdb", "dbac", "aada", "cdcda", "cdc", "dbc", "dbcb", "bdb", "ddbdd", "cadaa", "ddbc", "babb"}))
	fmt.Println(leetcode.WordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
	// output:
	// true
	// true
	// true
	// true
	// false
}

func ExampleTest() {
	s := "aaaaab"
	k := strings.Replace(s, "aaa", "", -1)
	fmt.Println(k)
	// output:
	// b
}
