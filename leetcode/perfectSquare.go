package leetcode

func PerfectSquare(n int) int {
	return numSquares(n)
}

func numSquares(n int) int {
	nums := []int{n}
	nexts := make([]int, 0)

	for i := 1 ; true ; i++ {
		for _, num := range nums {
			for j := 1 ; num - (j*j) >= 0 ; j++ {
				if num - (j*j) == 0 {
					return i
				}
				nexts = append(nexts, num - (j*j))
			}
		}
		nums = make([]int, len(nexts))
		copy(nums, nexts)
		nexts = make([]int, 0)
	}
	return n
}
