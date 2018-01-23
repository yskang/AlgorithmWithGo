package leetcode

func Exist(board [][]byte, word string) bool {
	return exist(board, word)
}

func exist(board [][]byte, word string) bool {
	visitBoard := make([][]bool, len(board))

	for i := 0; i < len(visitBoard); i++ {
		visitBoard[i] = make([]bool, len(board[0]))
	}

	for y, row := range board {
		for x, c := range row {
			if word[0] == c {
				if findFullWord(board, visitBoard, x, y, word) {
					return true
				}
			}
		}
	}
	return false
}

func findFullWord(board [][]byte, visitBoard [][]bool, x, y int, word string) bool {
	isFound := false
	if len(word) == 1 {
		return true
	}
	findChar(board, visitBoard, x, y, word, 0, &isFound)
	return isFound
}

func findChar(board [][]byte, visitBoard [][]bool, x, y int, word string, index int, isFound *bool) {
	if *isFound {
		return
	}
	visitBoard[y][x] = true

	if x+1 < len(board[0]) && board[y][x+1] == word[index+1] && !visitBoard[y][x+1] {
		if *isFound {
			return
		}
		if index+1 == len(word)-1 {
			*isFound = true
			return
		}
		findChar(board, visitBoard, x+1, y, word, index+1, isFound)
	}
	if y+1 < len(board) && board[y+1][x] == word[index+1] && !visitBoard[y+1][x] {
		if *isFound {
			return
		}
		if index+1 == len(word)-1 {
			*isFound = true
			return
		}
		findChar(board, visitBoard, x, y+1, word, index+1, isFound)
	}
	if x-1 >= 0 && board[y][x-1] == word[index+1] && !visitBoard[y][x-1] {
		if *isFound {
			return
		}
		if index+1 == len(word)-1 {
			*isFound = true
			return
		}
		findChar(board, visitBoard, x-1, y, word, index+1, isFound)
	}
	if y-1 >= 0 && board[y-1][x] == word[index+1] && !visitBoard[y-1][x] {
		if *isFound {
			return
		}
		if index+1 == len(word)-1 {
			*isFound = true
			return
		}
		findChar(board, visitBoard, x, y-1, word, index+1, isFound)
	}
	visitBoard[y][x] = false
}
