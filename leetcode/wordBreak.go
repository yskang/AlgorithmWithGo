package leetcode

// WordBreak is a solution of the problem "139. Word Break" in leetcode
func WordBreak(s string, wordDict []string) bool {
	return wordBreak(s, wordDict)
}

func wordBreak(s string, wordDict []string) bool {
	ans := false
	wordCheck(s, wordDict, &ans)
	return ans
}

func wordCheck(s string, wordDict []string, ans *bool) {
	if *ans == true {
		return
	}
	for _, word := range wordDict {
		if len(s) >= len(word) && s[:len(word)] == word {
			temp := s
			for {
				temp = temp[len(word):]
				if len(temp) >= len(word) && temp[:len(word)] == word {
					continue
				} else {
					break
				}
			}

			if len(temp) == 0 {
				*ans = true
				return
			}
			wordCheck(temp, wordDict, ans)
		}
	}
}
