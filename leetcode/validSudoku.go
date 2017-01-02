package main

import "fmt"

func main() {
	//["....5..1.",".4.3.....",".....3..1","8......2.","..2.7....","..2.7....",".....2...",".2.9.....","..4......"]

	m := [][]byte{[]byte("....5..1."),
		[]byte(".4.3....."),
		[]byte(".....3..1"),
		[]byte("8......2."),
		[]byte("..2.7...."),
		[]byte(".15......"),
		[]byte(".....2..."),
		[]byte(".2.9....."),
		[]byte("..4......")}

	fmt.Println(isValidSudoku(m))
}

func isValidSudoku(board [][]byte) bool {

	columnMap := make([]map[byte]int, 0)
	sectionMap := make([]map[byte]int, 0)

	for i := 0 ; i < 9 ; i++ {
		columnMap = append(columnMap, make(map[byte]int))
		sectionMap = append(sectionMap, make(map[byte]int))
	}

	for rowIndex, row := range(board) {
		rowMap := make(map[byte]int)
		for columnIndex, cell := range(row) {
			if _, exist := rowMap[cell]; exist {
				if cell != '.' {
					return false
				}
			}

			if _, exist := columnMap[columnIndex][cell]; exist {
				if cell != '.' {
					return false
				}
			}

			if _, exist := sectionMap[(rowIndex/3)*3+(columnIndex/3)][cell] ; exist {
				if cell != '.' {
					return false
				}
			}

			rowMap[cell] += 1
			columnMap[columnIndex][cell] += 1
			sectionMap[(rowIndex/3)*3+(columnIndex/3)][cell] += 1
		}
	}
	return true
}
