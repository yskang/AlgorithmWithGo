package main

import (
	"fmt"
)

func main() {
	getFunc := getCurve()

	for i := 0 ; i < 27 ; i++ {
		getFunc()
	}

	minusInCurve := *getFunc()

	var c int
	fmt.Scanf("%d\n", &c)

	for i := 0; i < c ; i++  {
		var n, p, l int
		fmt.Scanf("%d %d %d\n", &n, &p, &l)

		for j := p ; j < p + l ; j ++ {
			switch {
			case j % 6 == 1:
				fmt.Print("F")
			case j % 6 == 2:
				fmt.Print("X")
			case j % 6 == 3:
				isMinus := false
				for _, minIndex := range minusInCurve {
					if minIndex == j {
						fmt.Print("-")
						isMinus = true
						break
					}
				}
				if !isMinus {
					fmt.Print("+")
				}
			case j % 6 == 4:
				fmt.Print("Y")
			case j % 6 == 5:
				fmt.Print("F")
			case j % 6 == 0:
				isMinus := false
				for _, minIndex := range minusInCurve {
					if minIndex == j {
						fmt.Print("-")
						isMinus = true
						break
					}
				}
				if !isMinus {
					fmt.Print("+")
				}
			}
		}
		fmt.Print("\n")
	}
}

func getCurve() func() *[]int {
	minusList := make([]int, 0)
	length := 5

	return func() *[]int {
		center := int(length / 2) + 1

		l := len(minusList)
		isCenterMinus := false
		for i := 0 ; i < l ; i++ {
			if minusList[i] != center {
				minusList = append(minusList, minusList[i] + length + 1)
			} else {
				isCenterMinus = true
			}
		}
		if !isCenterMinus {
			minusList = append(minusList, center + length + 1)
		}


		length = length * 2 + 1

		return &minusList
	}
}
