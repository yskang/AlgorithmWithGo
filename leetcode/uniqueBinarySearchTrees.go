package leetcode

func NumTrees(n int) int {
	return numTrees(n)
}

func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	numTree := make([]int, n+1)

	numTree[0] = 1
	numTree[1] = 1
	numTree[2] = 2
	numTree[3] = 5

	for i := 4 ; i <= n ; i++ {
		sum := 0
		for j := 0 ; j < i ; j++ {
			 sum += numTree[j-0] * numTree[i-j-1]
		}
		numTree[i] = sum
	}

	return numTree[n]
}