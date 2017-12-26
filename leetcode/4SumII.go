package leetcode

// FourSumCount is solution of leetcode problem, 4 Sum II
func FourSumCount(A []int, B []int, C []int, D []int) int {
	return fourSumCount(A, B, C, D)
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	mapA := make(map[int]int)
	mapB := make(map[int]int)

	for _, a := range A {
		for _, b := range B {
			mapA[a+b]++
		}
	}

	for _, c := range C {
		for _, d := range D {
			mapB[c+d]++
		}
	}

	sum := 0
	for num, count := range mapA {
		if val, isExist := mapB[-num]; isExist {
			sum += val * count
		}
	}

	return sum
}
