package leetcode

import "fmt"

func BattleShipsInABoard() {
	fmt.Println(countBattleships([][]byte{{'X','.','.','X'},
		{'.','.','.','X'},
		{'.','.','.','X'}}))
	fmt.Println(countBattleships([][]byte{  {'X','.','.','X'},
						{'.','X','.','X'},
						{'.','X','.','.'},
						{'.','.','.','X'},
						{'X','X','.','X'}}))
}

func countBattleships(board [][]byte) int {
	count := 0

	for y := 0 ; y < len(board) ; y++ {
		for x := 0 ; x < len(board[0]) ; x++ {
			fmt.Print(string(board[y][x]))
			count += checkLeftAndUp(x, y, board)
		}
		fmt.Println();
	}

	return count
}
func checkLeftAndUp(x int, y int, board [][]byte) int {
	hasLeft := false
	hasUp := false

	if x != 0 && board[y][x-1] == 'X' {
		hasLeft = true
	}
	if y != 0 && board[y-1][x] == 'X' {
		hasUp = true
	}

	if board[y][x] == 'X' && hasLeft == false && hasUp == false {
		return 1
	}
	return 0
}