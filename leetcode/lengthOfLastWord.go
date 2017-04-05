package leetcode

import (
	"fmt"
	"strings"
)

func LengthOfLastWord() {
	fmt.Println(lengthOfLastWord("Hello World  "))
}

func lengthOfLastWord(s string) int {
	temp := strings.Split(strings.Trim(s, " "), " ")
	return len(temp[len(temp) - 1])
}
