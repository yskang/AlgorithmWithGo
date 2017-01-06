package main

import "fmt"

func main() {
	fmt.Println(findNthDigit(1))
	fmt.Println(findNthDigit(9))
	fmt.Println(findNthDigit(11))
	fmt.Println(findNthDigit(21))
	fmt.Println(findNthDigit(29))
	fmt.Println(findNthDigit(30))
	fmt.Println(findNthDigit(181))
	fmt.Println(findNthDigit(1189))
	fmt.Println(findNthDigit(1111))
	fmt.Println(findNthDigit(999))
	fmt.Println(findNthDigit(666))
}

func findNthDigit(n int) int {
	if n <= 9 {
		return n
	} else if n <= 9 + 10 * 2 * 9 {
		n -= 9

		if n % 2 == 1 {
			n = n / 20 + 1
		} else {
			n = (n % 21) / 2 - 1
		}
	} else if n <= 1189 {
		n -= 189

		if n % 3 == 1 {
			n = n / 300 + 1
		} else if n % 3 == 2 {
			n = n / 30 + 1
		} else {
			n = (n % 31) / 3 - 1
		}
	} else if n <= 11189 {

	} else if n <= 111189 {

	} else if n <= 1111189 {

	}

	return n
}