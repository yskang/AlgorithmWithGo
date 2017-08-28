package leetcode

func IsSubsequence(s string, t string) bool {
	return isSubsequence(s, t)
}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	nextBase := nextLetter(s)
	nextCompare := nextLetter(t)
	base := nextBase()
	compare := nextCompare()
	for {
		if compare == base {
			base = nextBase()
			if base == 0 {
				return true
			}
		}
		compare = nextCompare()
		if compare == 0 {
			break
		}
	}

	return false
}

func nextLetter(s string) func() rune {
	i := 0
	return func() rune {
		if i < len(s) {
			r := s[i]
			i += 1
			return rune(r)
		}
		return 0
	}
}
