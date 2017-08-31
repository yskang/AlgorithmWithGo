package leetcode

func MinSteps(n int) int {
	return minSteps(n)
}

func minSteps(n int) int {
	return minCost(n)
}

func minCost(n int) int {
	if n == 1 {
		return 0
	} else if n == 2{
		return 2
	} else if n == 3 {
		return 3
	}
	k := getMinDivider(n)

	if n == k {
		return n
	}

	return min(min(n, minCost(k) + n/k), min(n, minCost(n/k) + k))
}

func getMinDivider(n int) int {
	for i := 2 ; i < n ; i++ {
		if n % i == 0 {
			return i
		}
	}
	return n
}


func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
