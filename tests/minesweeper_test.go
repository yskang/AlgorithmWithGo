package tests

import (
	"testing"
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func TestClickMine(t *testing.T) {
	board := [][]byte{{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'}}

	expected := [][]byte{{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'X', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'}}

	click := []int{1,2}
	afterBoard := leetcode.Minesweeper(board, click)
	printBoard(afterBoard)
	if !isEqualBoard(expected, afterBoard) {
		t.Error("Wrong answer!")
		fmt.Println("expected is")
		printBoard(expected)
	}
}

func TestEmptySquareWithNoAdjacentMine(t *testing.T) {
	board := [][]byte{{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'}}

	expected := [][]byte{{'B', '1', 'E', '1', 'B'},
			     {'B', '1', 'M', '1', 'B'},
			     {'B', '1', '1', '1', 'B'},
			     {'B', 'B', 'B', 'B', 'B'}}

	click := []int{3,0}
	afterBoard := leetcode.Minesweeper(board, click)
	printBoard(afterBoard)
	if !isEqualBoard(expected, afterBoard) {
		t.Error("Wrong answer!")
		fmt.Println("expected is")
		printBoard(expected)
	}
}

func TestEmptySquareWithOneAdjacentMine(t *testing.T) {
	board := [][]byte{{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'}}

	expected := [][]byte{{'E', 'E', 'E', 'E', 'E'},
		             {'E', '1', 'M', 'E', 'E'},
			     {'E', 'E', 'E', 'E', 'E'},
			     {'E', 'E', 'E', 'E', 'E'}}

	click := []int{1,1}
	afterBoard := leetcode.Minesweeper(board, click)
	printBoard(afterBoard)
	if !isEqualBoard(expected, afterBoard) {
		t.Error("Wrong answer!")
		fmt.Println("expected is")
		printBoard(expected)
	}
}

func isEqualBoard(boardA [][]byte, boardB [][]byte) bool {
	if len(boardA) != len(boardB) || len(boardA[0]) != len(boardB[0]) {
		return false
	}

	for i := 0 ; i < len(boardA) ; i++ {
		for j := 0 ; j < len(boardA[0]) ; j++ {
			if boardA[i][j] != boardB[i][j] {
				return false
			}
		}
	}

	return true
}

func printBoard(board [][]byte) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(string(cell) + " ")
		}
		fmt.Println("")
	}
}
