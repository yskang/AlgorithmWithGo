package main

import (
	"fmt"
)

func main() {
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
				fmt.Print(getSymbol(j, n))
			case j % 6 == 4:
				fmt.Print("Y")
			case j % 6 == 5:
				fmt.Print("F")
			case j % 6 == 0:
				fmt.Print(getSymbol(j, n))
			}
		}
		fmt.Print("\n")
	}
}

func getSymbol(index int, nThGen int) string {
	if nThGen == 1 {
		return "+"
	}

	length := 2

	for i := 0 ; i < nThGen ; i++ {
		length = length * 2 + 1
	}

	symbol := ""
	indexParent := index % ((length + 1) / 2)
	if indexParent == 0 {
		symbol = "+"
	} else if indexParent == (length + 1) / 4 && index > ((length + 1) / 2) {
		symbol = reverse(getSymbol(indexParent, nThGen - 1))
	} else {
		symbol = getSymbol(indexParent, nThGen - 1)
	}

	return symbol
}

func reverse(symbol string) string {
	if symbol == "+" {
		return "-"
	}
	return "+"
}