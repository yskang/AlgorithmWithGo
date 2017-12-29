package leetcode

// NumIslands is a solution of the problem "200. Number of Islands" of leetcode
func NumIslands(grid [][]byte) int {
	return numIslands(grid)
}

func numIslands(grid [][]byte) int {
	count := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '1' {
				makeLandToSea(x, y, &grid)
				count++
			}
		}
	}

	return count
}

func makeLandToSea(x, y int, grid *[][]byte) {
	if x-1 >= 0 && (*grid)[y][x-1] == '1' {
		(*grid)[y][x-1] = '0'
		makeLandToSea(x-1, y, grid)
	}
	if x+1 < len((*grid)[0]) && (*grid)[y][x+1] == '1' {
		(*grid)[y][x+1] = '0'
		makeLandToSea(x+1, y, grid)
	}
	if y-1 >= 0 && (*grid)[y-1][x] == '1' {
		(*grid)[y-1][x] = '0'
		makeLandToSea(x, y-1, grid)
	}
	if y+1 < len((*grid)) && (*grid)[y+1][x] == '1' {
		(*grid)[y+1][x] = '0'
		makeLandToSea(x, y+1, grid)
	}
}
