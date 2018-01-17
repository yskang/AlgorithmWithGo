package leetcode

func MaximalSquare(matrix [][]byte) int {
	return maximalSquareDP(matrix)
}

func maximalSquareDP(matrix [][]byte) int {
	dp := make([][]int, len(matrix)+1)
	dp[0] = make([]int, len(matrix[0])+1)

	maxSquare := 0
	for y := 1; y <= len(matrix); y++ {
		dp[y] = make([]int, len(matrix[0])+1)
		for x := 1; x <= len(matrix[0]); x++ {
			if matrix[y-1][x-1] == '1' {
				dp[y][x] = minSq(dp[y][x-1], dp[y-1][x-1], dp[y-1][x]) + 1
				if maxSquare < dp[y][x] {
					maxSquare = dp[y][x]
				}
			}
		}
	}
	return maxSquare * maxSquare
}

func minSq(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	if c <= a && c <= b {
		return c
	}
	return 0
}

// my original solution
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
