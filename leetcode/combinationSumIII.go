package leetcode

import "fmt"

func CombinationSum3(k int, n int) [][]int {
	return combinationSum3(k, n)
}

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	nextComb := getNextComb(k)
	for {
		comb, hasNext := nextComb()
		if sumOf(comb) == n {
			ans = append(ans, comb)
		}
		if !hasNext {
			break
		}
	}

	return ans
}

func sumOf(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}

func getNextComb(size int) func()([]int, bool) {
	combination(make([]int, 5), 0, 5, 3, 0)
	n := 0
	hasNext := true
	return func() ([]int, bool) {

		n += 1
		if n == 10 {
			hasNext = false
		}
		return []int{}, hasNext
	}
}

func combination(arr []int, index int, n int, r int, target int) {
	if r == 0 {
		fmt.Println(arr, index)
	} else if target == n {
		return
	} else {
		arr[index] = target
		combination(arr, index+1, n, r-1, target+1)
		combination(arr, index, n, r, target+1)
	}
}
