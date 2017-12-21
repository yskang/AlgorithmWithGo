package leetcode

func GenerateParenthesis(n int) []string {
	return generateParenthesis(n)
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	genParenthesis(&ans, n, 0, "")
	return ans
}

func genParenthesis(ansList *[]string, toOpen int, toClose int, stringToMake string) {
	if toOpen == 0 && toClose == 0 {
		*ansList = append(*ansList, stringToMake)
		return
	}

	if toOpen > 0 {
		genParenthesis(ansList, toOpen-1, toClose+1, stringToMake+ "(")
	}

	if toClose > 0 {
		genParenthesis(ansList, toOpen, toClose-1, stringToMake+ ")")
	}
}


