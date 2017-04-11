package leetcode

func Minesweeper(board [][]byte, click []int) [][]byte {
	return updateBoard(board, click)
}
/*
Rules
1. If a mine ('M') is revealed, then the game is over - change it to 'X'.
2. If an empty square ('E') with no adjacent mines is revealed, then change it to revealed blank ('B') and all of its adjacent unrevealed squares should be revealed recursively.
3. If an empty square ('E') with at least one adjacent mine is revealed, then change it to a digit ('1' to '8') representing the number of adjacent mines.
4. Return the board when no more squares will be revealed.
*/
func updateBoard(board [][]byte, click []int) [][]byte{
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	if board[click[0]][click[1]] == 'E'{
		if numOfAdjMines := numOfAdjacentMine(board, click); numOfAdjMines == 0 {
			board[click[0]][click[1]] = 'B'
		 	revealEmptyBlocks(board, click)
	 	} else {
		 	board[click[0]][click[1]] = byte('0') + byte(numOfAdjMines)
	 	}
	}

	return board
}

func revealEmptyBlocks(board [][]byte, click []int) {
	for _, adj := range getAdjPositions(board, click) {
		if board[adj[0]][adj[1]] == 'E' {
			if numOfAdjMines := numOfAdjacentMine(board, adj); numOfAdjMines == 0 {
				board[adj[0]][adj[1]] = 'B'
				revealEmptyBlocks(board, adj)
			} else {
				board[adj[0]][adj[1]] = byte('0') + byte(numOfAdjMines)
			}
		}
	}
}

func numOfAdjacentMine(board [][]byte, click []int) int {
	numOfAdjMines := 0
	for _, adj := range getAdjPositions(board, click) {
		if board[adj[0]][adj[1]] == 'M' {
			numOfAdjMines += 1
		}
	}
	return numOfAdjMines
}

func getAdjPositions(board [][]byte, click []int) [][]int {
	adjs := make([][]int, 0)

	xs := []int{click[1]-1, click[1], click[1]+1}
	ys := []int{click[0]-1, click[0], click[0]+1}

	for x := 0 ; x < 3 ; x++ {
		for y := 0 ; y < 3 ; y++ {
			if xs[x] >= 0 && ys[y] >= 0 && xs[x] < len(board[0]) && ys[y] < len(board) && !(x == 1 && y == 1){
				adjs = append(adjs, []int{ys[y], xs[x]})
			}
		}
	}
	return adjs
}


