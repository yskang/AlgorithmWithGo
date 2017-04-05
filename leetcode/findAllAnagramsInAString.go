package leetcode

import (
	"fmt"
)

func FindAnagrams() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
func findAnagrams(s string, p string) []int {
	if s == "" {
		return nil
	}

	if len(s) < len(p) {
		return  nil
	}

	result := []int{}
	dict := make(map[rune]int)
	window := make(map[rune]int)

	for _, r := range(p) {
		dict[r] += 1
	}

	for _, r := range(s[0:len(p)]) {
		window[r] += 1
	}

	valid := true
	for key, _ := range(window) {
		if window[key] != dict[key] {
			valid = false
		}
	}
	if valid {
		result = append(result, 0)
	}
	//fmt.Println(window, dict)

	for start := 1 ; start <= len(s) - len(p); start++{
		window[rune(s[start-1])] -= 1
		window[rune(s[start+len(p)-1])] += 1

		//fmt.Println(window, dict)

		valid = true
		for key, _ := range(window) {
			if window[key] != dict[key] {
				valid = false
			}
		}
		if valid {
			result = append(result, start)
		}
	}

	return result
}


