package leetcode

import "log"

func MaximalSquare(matrix [][]byte) int {
	return maximalSquare(matrix)
}

func maximalSquare(matrix [][]byte) int {
	newMap := make([][]int, len(matrix))
	max := 0
	for y, row := range matrix {
		if row[0] == '1' {
			newMap[y] = append(newMap[y], 1)
		} else {
			newMap[y] = append(newMap[y], 0)
		}
		for _, column := range row[1:] {
			if column == '0' {
				newMap[y] = append(newMap[y], 0)
			} else {
				newMap[y] = append(newMap[y], newMap[y][len(newMap[y])-1]+1)
			}
		}
	}

	for _, row := range newMap {
		log.Println(row)
	}

	for y, row := range newMap {
		for x, v := range row {
			if v != 0 {
				log.Println("check ", x, y)
				m := getMaxCol(newMap, x, y)
				log.Println("down ", m)
				if max < m {
					max = m
				}
			}
		}
	}

	return max * max
}

func getMaxCol(matrix [][]int, x, y int) int {
	val := matrix[y][x]
	step := 0

	for v := val; v >= 1; v-- {
		log.Println("for val", v)
		for i := 1; y+i < len(matrix); i++ {
			if matrix[y+i][x] < v {
				break
			}
			step = i
		}
		if step+1 == v {
			return v
		}
	}

	return 1
}
