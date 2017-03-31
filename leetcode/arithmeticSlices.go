package main

import "fmt"

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1,2,3,4,5,6}))
	fmt.Println(numberOfArithmeticSlices([]int{}))
	fmt.Println(numberOfArithmeticSlices([]int{1,2}))
}

func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}

	totalCount := 0

	diffA := make([]int, 0)
	prevA := A[0]
	for _, a := range A {
		diffA = append(diffA, a - prevA)
		prevA = a
	}
	diffA = diffA[1:]

	count := 0
	prevDiff := diffA[0]
	for i := 0 ; i < len(diffA) ; i++ {
		if diffA[i] == prevDiff {
			count += 1
		} else {
			totalCount += int((count) * (count - 1) / 2)
			count = 1
			prevDiff = diffA[i]
		}
	}
	totalCount += int((count) * (count - 1) / 2)

	return totalCount
}