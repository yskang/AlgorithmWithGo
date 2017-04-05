package leetcode

import (
	"fmt"
	"sort"
)

func CountBits() {
	a := []string{"11110", "11111", "10111"}
	sort.Strings(a)
	fmt.Println(a)

	//fmt.Println(countBits(16))
}

func countBits(num int) []int {
	result := make([]int, 0)
	result = append(result, []int{0, 1}...)
	width := 1

	for {
		part := result[len(result)-width:]
		temp := make([]int, len(part) * 2)
		for i := 0 ; i < len(part) ; i++ {
			temp[i] = part[i]
		}

		for i := len(part) ; i < len(part) * 2 ; i++ {
			temp[i] = part[i-len(part)] + 1
		}

		result = append(result, temp...)
		if len(result) > num {
			return result[:num+1]
		}

		width *= 2
	}
}
