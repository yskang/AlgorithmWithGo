package leetcode

import "log"

func FindDiagonalOrder(matrix [][]int) []int {
	return findDiagonalOrder(matrix)
}

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	const (
		UP = true
		DOWN = false
	)

	width, height := len(matrix[0]), len(matrix)
	x, y := 0, 0
	index := 0
	ans := make([]int, width*height)
	direction := UP

	for index < len(ans) {
		log.Println("try", x, y, direction)
		ans[index] = matrix[y][x]
		log.Println(ans)

		if direction {
			x += 1
			y -= 1
			if x >= width {
				x -= 1
				y += 2
				direction = DOWN
			} else if y < 0 {
				y += 1
				direction = DOWN
			}
		} else {
			x -= 1
			y += 1
			if y >= height {
				x += 2
				y -= 1
				direction = UP
			} else if x < 0 {
				x += 1
				direction = UP
			}
		}

		index += 1
	}

	return ans
}
