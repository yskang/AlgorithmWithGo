// https://www.acmicpc.net/problem/2747
// https://www.acmicpc.net/problem/2748

package main

import "fmt"

func main() {
	num := 0
	fmt.Scanf("%d\n", &num)

	f := getFibonacci()

	fibonacciNum := uint64(0)

	for i := 0 ; i < num ; i++ {
		fibonacciNum = f()
	}

	fmt.Println(fibonacciNum)
}

func getFibonacci() func() uint64 {
	fst, snd := uint64(0), uint64(1)

	return func() uint64 {
		fst, snd = snd, fst + snd
		return fst
	}
}
