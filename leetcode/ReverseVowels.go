package leetcode

import (
	"fmt"
)

func ReverseVowels() {
	fmt.Println(reverseVowels("Hello World"))
	fmt.Println(reverseVowels("aA"))
}

func reverseVowels(s string) string {

	sBytes := []byte(s)

	vowels := make(map[byte]bool)
	vowels['a'] = true
	vowels['e'] = true
	vowels['i'] = true
	vowels['o'] = true
	vowels['u'] = true
	vowels['A'] = true
	vowels['E'] = true
	vowels['I'] = true
	vowels['O'] = true
	vowels['U'] = true

	backStart := len(s) - 1

	for i := 0 ; i < len(s) ; i++ {
		if vowels[s[i]] {
			temp := sBytes[i]
			for j := backStart ; j >= 0 ; j-- {
				if i == j || i > j {
					return string(sBytes)
				}
				if vowels[sBytes[j]] {
					sBytes[i] = s[j]
					sBytes[j] = temp
					backStart = j - 1
					break
				}
			}
		}
	}
	return s
}
