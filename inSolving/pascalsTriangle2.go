package main

import "fmt"

func main() {
	fmt.Println(getRow(10))
}


func getRow(rowIndex int) []int {
	cache := make(map[int][]int)

	return getPascalTriangleRow(rowIndex, cache)
}

func getPascalTriangleRow(rowIndex int, cache map[int][]int) []int {

	if value, isExist := cache[rowIndex]; isExist {
		return value
	}

	if rowIndex == 0 {
		cache[rowIndex] = []int{1}
		return cache[rowIndex]
	}

	row := []int{1}

	for i := 1 ; i < rowIndex ; i++ {
		row = append(row, getPascalTriangleRow(rowIndex - 1, cache)[i] + getPascalTriangleRow(rowIndex - 1, cache)[i-1])
	}
	row = append(row, 1)

	cache[rowIndex] = row
	return row
}
