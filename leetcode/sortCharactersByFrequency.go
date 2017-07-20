package leetcode

import (
	"sort"
	"strings"
)

func SortCharactersByFrequency(s string) string {
	return frequencySort(s)
}

func frequencySort(s string) string {
	freqMap := make(map[rune]int)

	for _, letter := range s {
		freqMap[letter] += 1
	}

	frequency := make([]freqLetter, 0)
	for letter, freq := range freqMap {
		frequency = append(frequency, freqLetter{letter, freq})
	}

	sort.Slice(frequency, func(i, j int) bool {
		return frequency[i].count > frequency[j].count
	})

	result := make([]string, 0)
	for _, elem := range frequency {
		for i := 0 ; i < elem.count ; i++ {
			result = append(result, string(elem.letter))
		}
	}

	return strings.Join(result, "")
}

type freqLetter struct {
	letter rune
	count int
}
