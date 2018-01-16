package leetcode

func MaximalSquare(matrix [][]byte) int {
	return maximalSquare(matrix)
}

func maximalSquare(matrix [][]byte) int {
	newMap := make([][]int, len(matrix))
	max := 0
	for y, row := range matrix {
		newMap[y] = make([]int, len(row))
		newMap[y][0] = int(row[0] - '0')
		for i, column := range row[1:] {
			if column == '0' {
				newMap[y][i+1] = 0
			} else {
				newMap[y][i+1] = newMap[y][i] + 1
			}
		}
	}

	for y, row := range newMap {
		for x, v := range row {
			if max < v {
				m := getMaxCol(newMap, x, y)
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
