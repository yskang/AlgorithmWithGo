package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countPrimes(50))
	fmt.Println(countPrimes(100))
	fmt.Println(countPrimes(500))
	fmt.Println(countPrimes(1500000))
	fmt.Println(countPrimes(499979))
}

func countPrimes(n int) int {
	pointer := 2
	buffer := [1500000]bool{}

	if n < 3 {
		return 0
	}

	for {
		buffer[pointer] = false

		for i := pointer + pointer ; i < n ; i += pointer {
			buffer[i] = true
		}

		if pointer > int(math.Sqrt(float64(n))) {
			break
		}

		for j := pointer + 1 ; j < n ; j++ {
			if buffer[j] == false {
				pointer = j
				break
			}
		}
	}

	count := 0

	for _, check := range(buffer[2:n]) {
		if check == false {
			count += 1
		}
	}

	return count
}
