package leetcode

import "strings"

func ReverseWords(s string) string {
	return reverseWords(s)
}

func reverseWords(s string) string {
	reverse := ""
	sentence := strings.Split(s, " ")

	for _, word := range sentence {
		reverse += reverseSingleWord(word) + " "
	}

	reverse = strings.TrimSpace(reverse)
	return reverse
}

func reverseSingleWord(s string) string {
	reverse := ""

	for i := len(s) - 1 ; i >=0 ; i-- {
		reverse += string(s[i])
	}
	return reverse
}