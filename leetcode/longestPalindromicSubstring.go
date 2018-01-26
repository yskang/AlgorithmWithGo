package leetcode

import (
	"fmt"
)

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

func longestPalindrome(s string) string {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}

	maxStr := ""

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] == s[len(s)-1-j] {
				dp[i+1][j+1] = dp[i][j] + 1
				if dp[i+1][j+1] > len(maxStr) && isPal(s[i+1-dp[i+1][j+1]:i+1]) {
					maxStr = s[i+1-dp[i+1][j+1] : i+1]
				}
			}
		}
	}

	for _, row := range dp {
		fmt.Println(row)
	}

	return maxStr
}

func longestPalindromeBruteForce(s string) string {
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
