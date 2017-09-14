package leetcode

import "fmt"

func NThUglyNumber(n int) int {
	return nthUglyNumber(n)
}

func nthUglyNumber(n int) int {
	result := make([]int, 0)
	result = append(result, []int{1}...)
	i, j, k := 0, 0, 0
	for len(result) < n {
		result = append(result, minUgly(result[i]*2, minUgly(result[j]*3, result[k]*5)))
		if result[len(result)-1] == result[i]*2 {
			i += 1
		}
		if result[len(result)-1] == result[j]*3 {
			j += 1
		}
		if result[len(result)-1] == result[k]*5 {
			k += 1
		}
	}
	fmt.Println(result)
	return result[len(result)-1]
}

func minUgly(a, b int) int {
	if a < b {
		return a
	}
	return b
}