package leetcodeTest

func GenerateParenthesis(n int) []string {
	return generateParenthesis(n)
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	addingpar(&res, "", n, 0)
	return res
}
func addingpar(v *[]string, str string, toOpen int, toClose int) {
	if toOpen == 0 && toClose == 0 {
		*v = append(*v, str)
		return
	}
	if toClose > 0 {
		addingpar(v, str + ")", toOpen, toClose-1)
	}
	if toOpen > 0 {
		addingpar(v, str + "(", toOpen-1, toClose+1)
	}
}

