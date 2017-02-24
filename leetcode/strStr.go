package main

import "fmt"

func main() {
	fmt.Println(strStr("a", ""))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for index, letter := range(haystack) {
		if index <= len(haystack) - len(needle) && letter == rune(needle[0]) {
			if haystack[index:index+len(needle)] == needle {
				return index
			}
		}
	}
	return -1
}