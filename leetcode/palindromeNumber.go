package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPalindrome(123454321))
}

func isPalindrome(x int) bool {


	i := 10
	l := 1
	for {
		if x % i == x {
			i = i / 10
			break
		}
		i *= 10
		l += 1
	}

	for k := 0 ; k < l/2 ; k++ {
		first := (x / (int(math.Pow10(l-1-k)))) % 10
		last := (x % (int(math.Pow10(k+1))))/int(math.Pow10(k))

		if first != last {
			return false
		}
	}


	return true
}
