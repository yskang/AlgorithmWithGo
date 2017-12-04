package leetcode

import (
	"fmt"
)

func CombinationSum3(k int, n int) [][]int {
	return combinationSum3(k, n)
}

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	next := make(chan int)

	for c := range combGen(9, k, next) {
		sum := 0
		for _, s := range c {
			sum += s + 1
		}
		if sum == n {
			temp := make([]int, k)
			for i := 0 ; i < k ; i++ {
				temp[i] = c[i]+1
			}
			ans = append(ans, temp)
		}
		next <- 0
	}

	return ans
}


func GenSum(a, b int) string {
	next:= make(chan int)
	for c := range combGen(5, 3, next) {
		fmt.Println(c)
		next <- 0
	}
	return ""
}

func combGen(a, b int, nc chan int) <-chan []int {
	c := make(chan []int)
	go func(n, m int) {
		s := make([]int, m)
		last := m - 1
		var rc func(int, int)

		rc = func(i, next int) {
			for j := next; j < n; j++ {
				s[i] = j
				if i == last {
					c <- s
					<- nc
					if s[0] == a - b {
						close(c)
						return
					}
				} else {
					rc(i+1, j+1)
				}
			}
			return
		}
		rc(0, 0)
	} (a, b)

	return c
}