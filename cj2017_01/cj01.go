package cj2017_01

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func CJ_2017_01() {
	args := os.Args[1:]
	inputFileName := args[0]
	outputFileName := strings.Replace(inputFileName, ".in", ".out", 1)

	results := make([]string, 0)

	inputs := readFile(inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0 ; t < T ; t++ {
		D, err := strconv.Atoi(inputs.Pop())
		checkErr(err)
		base := inputs.Pop()
		anagram := inputs.Pop()
		//fmt.Println(D, base, anagram)

		results = append(results, checkAnagram(D, base, anagram))
	}

	result := strings.Join(results, "\n")

	writeResultFile(outputFileName, []byte(result))
}
func checkAnagram(D int, base string, anagram string) string {
	if len(base) != len(anagram) {
		return "X"
	}

	baseMap := make(map[byte][]int)
	anagramMap := make(map[byte][]int)

	for i := 0 ; i < len(base) ; i++ {
		baseMap[base[i]] = append(baseMap[base[i]], i)
		anagramMap[anagram[i]] = append(anagramMap[anagram[i]], i)
	}

	for key, value := range baseMap {
		if val, isExist := anagramMap[key]; isExist {
			if len(val) != len(value) {
				return "X"
			}
		} else {
			return "X"
		}
	}

	for key, value := range baseMap {
		for i := 0 ; i < len(value) ; i++ {
			var diff int
			if value[i] > anagramMap[key][i] {
				diff = value[i] - anagramMap[key][i]
			} else {
				diff = anagramMap[key][i] - value[i]
			}
			if diff > D {
				return "X"
			}
		}
	}

	return "O"
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