// https://www.acmicpc.net/problem/2749
// success!

package baekjoon

import "fmt"

func FibonacciPeriod() {
	mod := uint64(1000000)
	p := uint64(mod/10 * 15)
	fibo := make([]uint64, p)
	fibo[0] = 0
	fibo[1] = 1

	var n uint64
	fmt.Scanf("%d\n", &n)

	for i := uint64(2) ; i < p ; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
		fibo[i] = fibo[i] % mod
	}

	index := uint64(n%p)

	fmt.Println(fibo[index])
}

// wrong. why????
//package main
//
//import "fmt"
//
//func main() {
//	var num int
//	fmt.Scanf("%d\n", &num)
//
//	fst, snd := uint64(0), uint64(1)
//
//	for i := 0 ; i < num % 1500000 ; i++ {
//		fst, snd = snd, (fst + snd) % 1000000
//	}
//
//	fmt.Println(fst)
//}
