package leetcode

import (
	"fmt"
)

func TrailingZeroes() {
	fmt.Println(trailingZeroes(300))
}

func trailingZeroes(n int) int {
	count := 0

	for i := 5 ; i <= n ; i*=5 {
		for j := 1 ; j * i <= n ; j++ {
			count += 1
		}
	}

	return count
}
