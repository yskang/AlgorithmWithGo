package leetcode


func IntegerBreak(n int) int {
	return integerBreak(n)
}

func integerBreak(n int) int {
	products := make([]int, n+1)
	products[1] = 1

	for i := 2 ; i <= n ; i ++ {
		for j := 1 ; j < i ; j++ {
			products[i] = max(products[i], max(j, products[j]) * max(i-j, products[i-j]))
		}
	}

	return products[n]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}