package leetcode

import (
	"fmt"
)

func BeautifulArrangement2() {
	fmt.Println(countArrangement2(5))
}

func countArrangement2(N int) int {
	count := 0
	count = checkCount(N, N, make([]bool, N+1), count)
	return count
}

func checkCount(N int, position int, visited []bool, count int) int{
	if position == 0 {
		count += 1
		return count
	}

	for i := N ; i > 0 ; i-- {
		if !visited[i] && (position % i == 0 || i % position == 0) {
			visited[i] = true
			count = checkCount(N, position - 1, visited, count)
			visited[i] = false
		}
	}

	return count
}

