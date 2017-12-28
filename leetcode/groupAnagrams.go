package leetcode

import (
	"strconv"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	return groupAnagrams(strs)
}

func groupAnagrams(strs []string) [][]string {
	dic := make(map[string][]string)
	for _, str := range strs {
		charMap := make(map[rune]int)
		for _, c := range str {
			charMap[c]++
		}
		key := make([]string, 0)
		for i := 'a'; i <= 'z'; i++ {
			carCount := charMap[rune(i)]
			if carCount != 0 {
				key = append(key, string(i)+strconv.Itoa(carCount))
			}
		}
		keyWord := strings.Join(key, "")
		dic[keyWord] = append(dic[keyWord], str)
	}

	ans := make([][]string, 0)
	for _, words := range dic {
		ans = append(ans, words)
	}

	return ans
}
