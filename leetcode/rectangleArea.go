package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(computeArea(-2, -2, 2, 2, -2, -2, 2, 2))
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	left := math.Max(float64(A), float64(E))
	right := math.Max(math.Min(float64(C), float64(G)), left)
	bottom := math.Max(float64(B), float64(F))
	top := math.Max(math.Min(float64(D), float64(H)), bottom)

	return int((float64(C)-float64(A)) * (float64(D)-float64(B)) - (right-left)*(top-bottom) + (float64(G)-float64(E))*(float64(H)-float64(F)))
}
