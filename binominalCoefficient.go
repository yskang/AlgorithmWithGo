// https://www.acmicpc.net/problem/11050

package main

import "fmt"

func main() {
	var n uint
	var k int

	fmt.Scanf("%d %d\n", &n, &k)

	var bc int

	if k >= 0 && k <= int(n) {
		bc = factorial(n) / (factorial(uint(k)) * factorial(n-uint(k)))
	} else if k < 0 || k > int(n) {
		bc = 0
	}

	fmt.Println(bc)
}

func factorial(i uint) int {
	if i == 0 {
		return 1
	}
	return int(i) * factorial(i-1)
}
