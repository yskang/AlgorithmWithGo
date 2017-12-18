package leetcode

import (
	"strings"
	"math"
)

func ShortestCompletingWord(licensePlate string, words []string) string {
	return shortestCompletingWord(licensePlate, words)
}

func shortestCompletingWord(licensePlate string, words []string) string {
	dic := make(map[rune]int)

	for _, c := range strings.ToLower(licensePlate) {
		if c >= 'a' && c <= 'z' {
			dic[c] += 1
		}
	}

	cans := make([]string, 0)
	shortestLength := math.MaxInt64

	for _, str := range words {
		if checkStr(str, dic) {
			cans = append(cans, str)
			if len(str) < shortestLength {
				shortestLength = len(str)
			}
		}
	}

	for _, str := range cans {
		if len(str) == shortestLength {
			return str
		}
	}

	return ""

}

func checkStr(s string, dic map[rune]int) bool {
	for char, count := range dic {
		 if strings.Count(strings.ToLower(s), string(char)) >= count {
		 	continue
		 } else {
		 	return false
		 }
	}
	return true
}

