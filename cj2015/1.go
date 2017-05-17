package cj2015

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func CJ_2015_1 () {
	pathName := "/home/yskang/Documents/2015_1/"
	inputFileName := "problem1.in"
	outputFileName := "result1.out"
	results := make([]string, 0)

	inputs := readFile(pathName + inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0 ; t < T ; t++ {

		bestScore := 10000
		bestPhoneNumber := ""
		N, err := strconv.Atoi(inputs.Pop())
		checkErr(err)
		for i := 0 ; i < N ; i++ {
			phoneNumber := inputs.Pop()
			score := getScore(phoneNumber)
			if score < bestScore {
				bestScore = score
				bestPhoneNumber = phoneNumber
			}
		}

		results = append(results, bestPhoneNumber)
	}

	result := strings.Join(results, "\n")

	writeResultFile(pathName + outputFileName, []byte(result))
}

func getScore(number string) int {
	score := 0
	prevNumber := number[0]
	prevDirection := ""
	for i := 1 ; i < len(number) ; i++ {
		if i == 1 {
			if prevNumber == number[i] {
				score += 0
			} else if bConnected, direction := isConnected(prevNumber, number[i]); bConnected {
				prevDirection = direction
				score += 1
			} else {
				prevDirection = direction
				score += 2
			}
		} else {
			if prevNumber == number[i] {
				score += 0
			} else if bConnected, direction := isConnected(prevNumber, number[i]); bConnected {
				if direction == prevDirection {
					score += 1
				} else {
					score += 2
				}
				prevDirection = direction
			} else {
				score += 3
			}
		}
		prevNumber = number[i]
	}
	return score
}

func isConnected(prevNumber byte, currentNumber byte) (bool, string) {
	switch prevNumber {
	case '0':
		switch currentNumber {
		case '7':
			return true, "UL"
		case '8':
			return true, "U"
		case '9':
			return true, "UR"
		default:
			return false, "NO"
		}
	case '1':
		switch currentNumber {
		case '2':
			return true, "R"
		case '4':
			return true, "D"
		case '5':
			return true, "DR"
		default:
			return false, "NO"
		}
	case '2':
		switch currentNumber {
		case '1':
			return true, "L"
		case '3':
			return true, "R"
		case '4':
			return true, "DL"
		case '5':
			return true, "D"
		case '6':
			return true, "DR"
		default:
			return false, "NO"
		}
	case '3':
		switch currentNumber {
		case '2':
			return true, "L"
		case '5':
			return true, "DL"
		case '6':
			return true, "D"
		default:
			return false, "NO"
		}
	case '4':
		switch currentNumber {
		case '1':
			return true, "U"
		case '2':
			return true, "UR"
		case '5':
			return true, "R"
		case '7':
			return true, "D"
		case '8':
			return true, "DR"
		default:
			return false, "NO"
		}
	case '5':
		switch currentNumber {
		case '0':
			return false, "NO"
		case '1':
			return true, "UL"
		case '2':
			return true, "U"
		case '3':
			return true, "UR"
		case '4':
			return true, "L"
		case '6':
			return true, "R"
		case '7':
			return true, "DL"
		case '8':
			return true, "D"
		case '9':
			return true, "DR"
		}
	case '6':
		switch currentNumber {
		case '2':
			return true, "UL"
		case '3':
			return true, "U"
		case '5':
			return true, "L"
		case '8':
			return true, "DL"
		case '9':
			return true, "D"
		default:
			return false, "NO"
		}
	case '7':
		switch currentNumber {
		case '4':
			return true, "U"
		case '5':
			return true, "UR"
		case '8':
			return true, "R"
		default:
			return false, "NO"
		}
	case '8':
		switch currentNumber {
		case '4':
			return true, "UL"
		case '5':
			return true, "U"
		case '6':
			return true, "UR"
		case '7':
			return true, "L"
		case '9':
			return true, "R"
		default:
			return false, "NO"
		}
	case '9':
		switch currentNumber {
		case '5':
			return true, "UL"
		case '6':
			return true, "U"
		case '8':
			return true, "L"
		default:
			return false, "NO"
		}
	}
	return false, "NO"
}

func readFile(path string) Queue {
	t1 := time.Now()
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(strings.Replace(string(dat), "\r", "", -1), "\n")
	size := len(lines)
	fmt.Println("read file time : ", time.Now().Sub(t1))
	return Queue{lines, size, 0, size-1, size-1}
}

func writeResultFile(path string, data []byte) {
	ioutil.WriteFile(path, data, os.ModePerm)
}

type Queue struct {
	nodes []string
	size  int
	head  int
	tail  int
	count int
}

func (q *Queue) Pop() string {
	if q.count == 0 {
		return "Empty"
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}