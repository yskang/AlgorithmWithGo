package cj2015_3

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func Cj2015_3() {
	pathName := "/home/yskang/Downloads/2014_03/"
	inputFileName := "problem0.in"
	outputFileName := "result0.out"
	results := make([]string, 0)

	inputs := readFile(pathName + inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0; t < T; t++ {
		line := strings.Split(inputs.Pop(), " ")
		N, err := strconv.Atoi(line[0])
		checkErr(err)
		M, err := strconv.Atoi(line[1])
		checkErr(err)

		fmt.Println(N, M)
		pairs := make(map[string][]string)

		for i := 0; i < M; i++ {
			pair := strings.Split(inputs.Pop(), " ")
			pairs[pair[0]] = append(pairs[pair[0]], pair[1])
			pairs[pair[1]] = append(pairs[pair[1]], pair[0])
		}

		fmt.Println(pairs)

		results = append(results, strconv.Itoa(t))
	}

	result := strings.Join(results, "\n")

	writeResultFile(pathName+outputFileName, []byte(result))
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
	return Queue{lines, size, 0, size - 1, size - 1}
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
