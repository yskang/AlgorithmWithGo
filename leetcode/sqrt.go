package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(9))
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	start, end := 1, x

	for {
		if start == end {
			return start
		} else if end - start == 1 {
			if end*end < x {
				return end
			}
			return start
		}

		mid := (start + end) / 2
		if p := mid * mid; p > x {
			end = mid
		} else if p < x {
			start = mid
		} else {
			return mid
		}

	}
}
