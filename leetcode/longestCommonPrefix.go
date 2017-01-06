package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"aa", "a"}))
	fmt.Println(longestCommonPrefix([]string{}))
}

func longestCommonPrefix(strs []string)	string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1 ; i < len(strs) ; i++ {
		length := 0

		if len(strs[i]) < len(prefix) {
			length = len(strs[i])
			prefix = prefix[:length]
		} else {
			length = len(prefix)
		}

		for j := 0 ; j < length ; j++ {
			if prefix[j:j+1] != strs[i][j:j+1] {
				prefix = prefix[:j]
				break
			}
		}

	}

	return prefix
}
