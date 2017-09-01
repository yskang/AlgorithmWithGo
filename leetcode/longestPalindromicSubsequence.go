package leetcode

func LongestPalindromeSubseq(s string) int {
	return longestPalindromeSubseq(s)
}

func longestPalindromeSubseq(s string) int {
	if len(s) == 1 {
		return 1
	}
	memo := make([][]int, len(s))
	for i := range s {
		memo[i] = make([]int, len(s))
	}

	for k := 0 ; k < len(s) ; k++ {
		for i := 0 ; i < len(s) ; i++ {
			if (i+k) < len(s) {
				memo[i][i+k] = lengthPal(i, i+k, s, memo)
			} else {
				break
			}
		}
	}

	return memo[0][len(s)-1]
}

func lengthPal(i int, j int, s string, memo [][]int) int {
	if j - i == 0 {
		return 1
	}

	if s[i] == s[j] {
		if j - i == 1 {
			return 2
		} else if j - i == 2 {
			return 3
		}
		return memo[i+1][j-1] +2
	}

	return maxP(memo[i+1][j], memo[i][j-1])
}

func maxP(a int, b int) int {
	if a > b {
		return a
	}
	return b
}