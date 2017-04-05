package leetcode

import (
	"fmt"
)

func ConvertToTitle() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(27))
	fmt.Println(convertToTitle(51))
	fmt.Println(convertToTitle(52))
	fmt.Println(convertToTitle(53))
	fmt.Println(convertToTitle(54))
	fmt.Println(convertToTitle(55))
	fmt.Println(convertToTitle(78))
	fmt.Println(convertToTitle(79))
}

func convertToTitle(n int) string {
	title := ""
	for n != 0 {
		digit := 'A' + (n - 1) % 26
		title = string(digit) + title
		n = (n - 1) / 26
	}

	return title
}