package leetcode

import "fmt"

func Subsets(nums []int) [][]int {
	return subsets(nums)
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	for i := 0 ; i <= len(nums) ; i++ {
		ans = append(ans, getSubset(nums, i)...)
	}
	return ans
}

func getSubset(nums []int, k int) [][]int {
	ans := make([][]int, 0)
	if k == 0 {
		return [][]int{nil}
	} else if k == len(nums) {
		return append([][]int{}, nums)
	}
	c := make(chan int)
	for comb := range combGen2(len(nums), k, c) {
		temp := make([]int, len(comb))
		for i, ci := range comb {
			temp[i] = nums[ci]
		}
		ans = append(ans, temp)
		c <- 0
	}
	fmt.Println(ans)
	return ans
}

func combGen2(a, b int, nc chan int) <-chan []int {
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