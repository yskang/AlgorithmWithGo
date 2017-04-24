package leetcode

func RotateImage(matrix [][]int) [][]int {
	rotateImage(matrix)
	return matrix
}

func rotateImage(matrix [][]int) {
	for i := len(matrix) / 2 - 1 ; i >= 0 ; i-- {
		opp := len(matrix) - 1 - i
		matrix[i], matrix[opp] = matrix[opp], matrix[i]
	}

	for i := 0 ; i < len(matrix) ; i++ {
		for j := i + 1 ; j < len(matrix[i]) ; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}

