package main

import (
	"fmt"
)

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
}

func isIsomorphic(s string, t string) bool {
	dic := make(map[rune]rune)
	dic2 := make(map[rune]rune)

	for index, c := range(s) {
		if dic[c] == 0 {
			dic[c] = rune(t[index])
		} else if dic[c] != rune(t[index]) {
			return false
		}
	}

	for index, c := range(t) {
		if dic2[c] == 0 {
			dic2[c] = rune(s[index])
		} else if dic2[c] != rune(s[index]) {
			return false
		}
	}

	return true
}
