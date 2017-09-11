package leetcode

func MinPathSum(grid [][]int) int {
	return minPathSum(grid)
}

func minPathSum(grid [][]int) int {
	memo := make(map[pos]int)
	return getMinPathSum(grid, len(grid[0])-1, len(grid)-1, memo)
}

func getMinPathSum(grid [][]int, x int, y int, memo map[pos]int) int {
	if val, isExist := memo[pos{x, y}]; isExist {
		return val
	}
	if x == 0 && y == 0 {
		return grid[0][0]
	} else if x == 0 {
		sum := 0
		for i := 0 ; i <= y ; i++ {
			sum += grid[i][0]
		}
		memo[pos{x, y}] = sum
		return sum
	} else if y == 0 {
		sum := 0
		for i := 0 ; i <= x ; i++ {
			sum += grid[0][i]
		}
		memo[pos{x, y}] = sum
		return sum
	}

	memo[pos{x, y}] = minAB(getMinPathSum(grid, x-1, y, memo), getMinPathSum(grid, x, y-1, memo)) + grid[y][x]
	return memo[pos{x, y}]
}

func minAB(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

type pos struct {
	x int
	y int
}