package leetcode

func CountSubstrings(s string) int {
	return countSubstrings(s)
}

func countSubstrings(s string) int {
	count := 0
	
	for i := 1 ; i <= len(s) ; i++ {
		count += numberOfPalindromesWithWindow(s, i)
	}	
	
	return count
}

func numberOfPalindromesWithWindow(s string, window int) int {
	start := 0
	count := 0
	for start + window <= len(s) {
		if checkPalindrom(s[start:start+window]) {
			count += 1
		}
		start += 1
	}
	return count
}

func checkPalindrom(s string) bool {
	for i := 0 ; i < len(s) / 2 ; i ++ {
		if s[i] != s[len(s) - 1 - i] {
			return false
		}
	}
	return true
}
