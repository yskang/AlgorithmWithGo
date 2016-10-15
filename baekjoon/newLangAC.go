// https://www.acmicpc.net/problem/5430

package baekjoon

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func printArray(array string, isReverse bool, start int, end int) {
	if start >= end {
		fmt.Println("[]")
		return
	}

	if !isReverse {
		fmt.Println("[" + array[start:end] + "]");
	} else {
		arrayStr := strings.Split(array[start:end], ",")
		result := "["
		for i := len(arrayStr) - 1; i >= 0; i-- {
			result = result + arrayStr[i] + ","
		}
		result = strings.TrimRight(result, ",")
		result = result + "]"

		fmt.Println(result)
	}
}

func NewLangAC() {
	numTest := 0
	fmt.Scanf("%d\n", &numTest)

	scanner := bufio.NewScanner(os.Stdin)

	inStrings := make([]string, 0)

	for scanner.Scan() {
		inStrings = append(inStrings, scanner.Text())

		if len(inStrings) == 3 {
			numTest = numTest - 1
			valid := true
			isReverse := false

			command := inStrings[0]
			arrayString := inStrings[2]

			start := 1
			end := len(arrayString) - 1

			for _, com := range command {
				switch com {
				case 'R':
					isReverse = !isReverse
				case 'D':
					if end - start <= 0 {
						valid = false
						break
					}
					if !isReverse {
						for j := start; true; j++ {
							if arrayString[j] == ',' || arrayString[j] == ']' {
								start = j + 1
								break
							}
						}
					} else {
						for j := end; true; j-- {
							if arrayString[j - 1] == ',' || arrayString[j - 1] == '[' {
								end = j - 1
								break
							}
						}
					}
				default:
					valid = false
					break
				}
			}

			if valid == false {
				fmt.Println("error")
			} else {
				printArray(arrayString, isReverse, start, end)
			}

			inStrings = inStrings[:0]
		}

		if numTest == 0 {
			break
		}
	}
}
