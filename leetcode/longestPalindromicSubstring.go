package leetcode

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

func longestPalindrome(s string) string {
	lgstPal := ""

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPal(s[i : j+1]) {
				if j-i+1 > len(lgstPal) {
					lgstPal = s[i : j+1]
				}
			}
		}
	}

	return lgstPal
}

func isPal(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i, c := range s[0 : len(s)/2] {
		if c != rune(s[len(s)-1-i]) {
			return false
		}
	}
	return true
}
