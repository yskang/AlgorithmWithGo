package leetcode

import (
	"strconv"
	"fmt"
)

func FractionAddition(expression string) string {
	return fractionAddition(expression)
}


func fractionAddition(expression string) string {
	numerator := 0
	denominator := 0

	sumNum := 0
	sumDenom := 1

	isNumerator := true

	for i := 0 ; i < len(expression) ; i++ {
		start := i

		if expression[i] == '-' {
			i += 1
		}

		for expression[i] >= '0' && expression[i] <= '9' {
			i += 1
			if i >= len(expression) {
				break
			}
		}

		if start != 0 && expression[start - 1] == '-' {
			start -= 1
		}

		num, _ := strconv.Atoi(expression[start:i])

		if isNumerator {
			numerator = num
		} else {
			denominator = num

			sumNum = sumNum* denominator + sumDenom* numerator
			sumDenom = sumDenom * denominator
		}

		isNumerator = !isNumerator
	}

	gcdNum := 1
	if sumNum < 0 {
		gcdNum = gcd(-1*sumNum, sumDenom)
	} else {
		gcdNum = gcd(sumNum, sumDenom)
	}

	return fmt.Sprintf("%d/%d", int(sumNum/gcdNum), int(sumDenom/gcdNum))
}

func gcd(a int, b int) int {
	temp := 0
	for b > 0 {
		temp = a
		a = b
		b = temp % b
	}
	return a
}