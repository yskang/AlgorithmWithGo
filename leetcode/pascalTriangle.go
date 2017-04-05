package leetcode

import "fmt"

func PascalTriangle() {
	fmt.Println(generate(1))
}

func generate(numRows int) [][]int {
	pascalTriangle := [][]int{[]int{1}, []int{1,1}}

	for i := 2 ; i < numRows ; i ++ {
		row := []int{1}

		for j := 0 ; j < len(pascalTriangle[i - 1]) - 1; j++ {
			row = append(row, pascalTriangle[i-1][j] + pascalTriangle[i-1][j+1])
		}

		row = append(row, 1)

		pascalTriangle = append(pascalTriangle, row)
	}

	return pascalTriangle[:numRows]
}
