// https://www.acmicpc.net/problem/2193

package baekjoon

import "fmt"

func PinaryNumber(n int) {
	var N int
	fmt.Scanf("%d\n", &N)

	type PNumber struct {
		zeros int
		ones int
	}

	pNumber := new(PNumber)
	pNumber.ones = 1
	pNumber.zeros = 0

	for i := 1 ; i < N ; i++ {
		pNumber.ones, pNumber.zeros = pNumber.zeros, pNumber.ones + pNumber.zeros
	}

	fmt.Printf("%d\n", pNumber.ones + pNumber.zeros)
}