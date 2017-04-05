package leetcode

import (
	"fmt"
	"strings"
)

func ValidPalindrome() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	sInLower := strings.ToLower(s)
	onlyAlphabet := ""
	reverseAlphabet := ""

	for _, letter := range(sInLower) {
		if (letter >= 'a' && letter <= 'z') || (letter >= '0' && letter <= '9') {
			onlyAlphabet += string(letter)
			reverseAlphabet = string(letter) + reverseAlphabet
		}
	}

	if onlyAlphabet == reverseAlphabet {
		return true
	}
	return false
}
