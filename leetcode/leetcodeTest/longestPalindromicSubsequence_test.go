package leetcodeTest

import (
	"fmt"
	"AlgorithmWithGo/leetcode"
)

func ExampleLongestPalindromeSubsequence() {
	fmt.Println(leetcode.LongestPalindromeSubseq("bbbab"))
	fmt.Println(leetcode.LongestPalindromeSubseq("cbbd"))
	fmt.Println(leetcode.LongestPalindromeSubseq("abcdef"))
	fmt.Println(leetcode.LongestPalindromeSubseq("aaa"))
	// output:
	// 4
	// 2
	// 1
	// 3
}

func ExampleLogestPalindromeLong() {
	fmt.Println(leetcode.LongestPalindromeSubseq("euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew"))
	// output:
	// 159
}