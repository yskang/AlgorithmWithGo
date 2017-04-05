package leetcode

import "fmt"

func ValidParentheses() {
	fmt.Println(isValid("]"))
}

func isValid(s string) bool {

	stack := make([]rune, 0)

	for _, char := range(s) {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			poped := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if char == ')' && poped == '('{
				continue
			} else if char == '}' && poped == '{'{
				continue
			} else if char == ']' && poped == '['{
				continue
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}