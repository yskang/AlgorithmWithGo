package main

import "fmt"

func main() {
	fmt.Println(arrangeCoins())
}

func arrangeCoins(n int) int {
	sum := 0
	i := 0

	for {
		sum = (1+i) * i / 2

		if sum > n {
			break
		}

		i += 1

	}

	return i-1
}
