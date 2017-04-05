package leetcode

import (
	"fmt"
	"strings"
	"strconv"
)

func ComplexNumberMultiply() {
	fmt.Printf(complexNumberMultiply("1+1i", "1+1i"))
}

func complexNumberMultiply(a string, b string) string {
	as := strings.Split(a, "+")
	bs := strings.Split(b, "+")

	aReal, _ := strconv.Atoi(as[0])
	aImage, _ := strconv.Atoi(strings.Replace(as[1], "i", "", 1))
	bReal, _ := strconv.Atoi(bs[0])
	bImage, _ := strconv.Atoi(strings.Replace(bs[1], "i", "", 1))

	real := strconv.Itoa(aReal * bReal - aImage * bImage)
	image := strconv.Itoa(aReal * bImage + aImage * bReal)

	return real + "+" + image + "i"
}