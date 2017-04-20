package leetcode

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}

func searchMatrix(matrix [][]int, target int) bool {

	totalLen := len(matrix) * len(matrix[0])

	start := 0
	end := totalLen - 1

	for start != end {
		center := int((start + end - 1)/2)
		if getElement(matrix, center) < target {
			start = center + 1
		} else {
			end = center
		}
	}

	return getElement(matrix, start) == target
}

func getElement(matrix [][]int, index int) int {
	rowNum := index / len(matrix[0])
	columnNum := index % len(matrix[0])
	return matrix[rowNum][columnNum]
}