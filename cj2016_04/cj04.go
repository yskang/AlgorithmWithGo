package cj2016_04

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"AlgorithmWithGo/myLibs"
)

func CJ_2016_04 () {
	pathName := "/home/yskang/Documents/cj_04/"
	inputFileName := "Set0.in"
	outputFileName := "Set0.out"
	results := make([]string, 0)

	inputs := readFile(pathName + inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0 ; t < T ; t++ {
		n := strings.Split(inputs.Pop(), " ")

		sum := "0"
		for i := 0 ;; i++ {
			sum = myLibs.BigAdd(sum, n[i % 5])

			if isRepeatNumber(sum) {
				break
			}
			if i == 100 {
				break
			}
		}

		fmt.Println(sum)

		results = append(results, strconv.Itoa(t))
	}

	result := strings.Join(results, "\n")

	writeResultFile(pathName + outputFileName, []byte(result))
}

func isRepeatNumber(number string) bool {
	fmt.Println(number)
	if strings.Count(number, string(number[0])) == len(number) {
		return true
	}
	return false
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