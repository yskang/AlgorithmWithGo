// https://www.acmicpc.net/problem/1149

package main

import (
	"fmt"
)

var valueTable = make([][]int, 0)

func main() {
	var numHouse, r, g, b int
	costTable := make([][]int, 0)

	fmt.Scanf("%d\n", &numHouse)

	for i := 0 ; i < numHouse; i++ {
		fmt.Scanf("%d %d %d\n", &r, &g, &b)
		valueTable = append(valueTable, []int{r, g, b})
	}

	costTable = append(costTable, valueTable[0])

	for i := 1 ; i < numHouse ; i++ {
		valueTable[i][0] = valueTable[i][0] + min(valueTable[i-1][1], valueTable[i-1][2])
		valueTable[i][1] = valueTable[i][1] + min(valueTable[i-1][0], valueTable[i-1][2])
		valueTable[i][2] = valueTable[i][2] + min(valueTable[i-1][0], valueTable[i-1][1])
	}

	minCost := 999999999
	for i := 0 ; i < 3 ; i++ {
		if minCost > valueTable[numHouse-1][i] {
			minCost = valueTable[numHouse-1][i]
		}
	}

	fmt.Println(minCost)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
