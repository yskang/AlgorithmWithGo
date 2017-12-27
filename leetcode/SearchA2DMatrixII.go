package leetcode

//SearchMatrixII is a solutin of search a 2D matrix II of leetcode
func SearchMatrixII(matrix [][]int, target int) bool {
	return searchMatrixII(matrix, target)
}

func searchMatrixII(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	x, y := 0, len(matrix)-1

	for {
		if x < 0 || x == len(matrix[0]) || y < 0 || y == len(matrix) {
			return false
		} else if matrix[y][x] == target {
			return true
		} else if matrix[y][x] > target {
			y--
		} else {
			x++
		}
	}
}

func searchMatrixII_(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	} else if len(matrix[0]) == 0 {
		return false
	}

	// find rows
	rows := make([]int, 0)
	for i, row := range matrix {
		if target >= row[0] && target <= row[len(row)-1] {
			rows = append(rows, i)
		}
	}
	// find columns
	columns := make([]int, 0)
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] <= target && matrix[len(matrix)-1][i] >= target {
			columns = append(columns, i)
		}
	}

	for _, y := range rows {
		for _, x := range columns {
			if matrix[y][x] == target {
				return true
			}
		}
	}

	return false
}
