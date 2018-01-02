package leetcode

// LetterCombinations is a solution of the probkem "17. Letter Combinations of a Phone Number" in leetcode
func LetterCombinations(digits string) []string {
	return letterCombinations(digits)
}

func letterCombinations(digits string) []string {
	getDigit := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}
	if digits == "" {
		return nil
	}
	strs := make([]string, 0)
	strs = append(strs, getDigit[digits[0]-'2']...)

	for _, digit := range digits[1:] {
		str := getDigit[digit-'2']
		temp := make([]string, len(strs))
		copy(temp, strs)
		strs = []string{}
		for _, s := range str {
			for _, t := range temp {
				t = t + s
				strs = append(strs, t)
			}
		}
	}

	return strs
}
