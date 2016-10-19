package main

import (
	"fmt"
)

func main() {
	dragonCurve := "FX"
	getFunc := getCurve()

	for i := 0 ; i < 27 ; i++ {
		getFunc()
	}

	dragonCurve = *getFunc()

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
				fmt.Print(string(dragonCurve[j/3-1]))
			case j % 6 == 4:
				fmt.Print("Y")
			case j % 6 == 5:
				fmt.Print("F")
			case j % 6 == 0:
				fmt.Print(string(dragonCurve[j/3-1]))
			}
		}
		fmt.Print("\n")
	}
}

func getCurve() func() *string {
	s1 := ""
	s2 := ""
	curve := "+"

	return func() *string {
		s1 = curve

		center := int(len(curve) / 2)
		if curve[center] == '+' {
			s2 = curve[:center] + "-" + curve[center+1:]
		} else {
			s2 = curve[:center] + "+" + curve[center+1:]
		}

		curve = s1 + "+" + s2
		return &curve
	}
}
