package leetcode

func CountNumbersWithUniqueDigits(n int) int {
	return countNumbersWithUniqueDigits(n)
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	numberOfUniqueDigits := make([]int, n + 1)

	numberOfUniqueDigits[1] = 10

	for i := 2 ; i <= n ; i++ {
		numberOfUniqueDigits[i] = 9
	}

	for i := 2 ; i <= n ; i++ {
		for multiplexer := 9 ; multiplexer > 10 - i ; multiplexer-- {
			numberOfUniqueDigits[i] *= multiplexer
		}
	}

	sumOfSmallerDigit := 0
	for i := 1 ; i < n ; i++ {
		sumOfSmallerDigit += numberOfUniqueDigits[i]
	}

	return sumOfSmallerDigit + numberOfUniqueDigits[n]
}