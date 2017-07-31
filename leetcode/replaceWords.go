package leetcode

import "strings"

func ReplaceWords(dict []string, sentence string) string {
	return replaceWords(dict, sentence)
}

func replaceWords(dict []string, sentence string) string {
	newWords := make([]string, 0)
	for index, word := range strings.Split(sentence, " ") {
		for _, root := range dict {
			if strings.HasPrefix(word, root) {
				newWords = append(newWords, root)
				break
			}
		}
		if len(newWords) == index {
			newWords = append(newWords, word)
		}
	}

	return strings.Join(newWords, " ")
}
