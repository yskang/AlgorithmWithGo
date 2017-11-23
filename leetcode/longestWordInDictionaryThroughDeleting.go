package leetcode

import (
	"sort"
)

func FindLongestWord(s string, d []string) string {
	return findLongestWord(s, d)
}

func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		if len(d[i]) > len(d[j]) {
			return true
		} else if len(d[i]) < len(d[j]) {
			return false
		}
		for k := 0 ; k < len(d[i]) ; k++ {
			if d[i][k] == d[j][k] {
				continue
			} else if d[i][k] < d[j][k] {
				return true
			} else {
				return false
			}
		}
		return false
	})

	for _, word := range d {
		i, j := 0, 0
		for i < len(s) {
			if word[j] != s[i] {
				i += 1
			} else {
				j += 1
				i += 1
			}
			if j == len(word) {
				return word
			}
		}
	}
	return ""
}