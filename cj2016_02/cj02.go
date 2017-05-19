package cj2016_02

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"sort"
)

func CJ_2016_02 () {
	pathName := "/home/yskang/Documents/cj_02/"
	inputFileName := "Set3.in"
	outputFileName := "Set3.out"
	results := make([]string, 0)

	inputs := readFile(pathName + inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0 ; t < T ; t++ {
		N, err := strconv.Atoi(inputs.Pop())
		checkErr(err)
		ps := make([]int, 0)

		for i := 0 ; i < N ; i++ {
			p, _ := strconv.Atoi(inputs.Pop())
			ps = append(ps, p)
		}

		sort.Ints(ps)
		fmt.Println(ps)
		maxDiff := 0

		for i := 0 ; i < len(ps) - 2 ; i++ {
			if ps[i+2] - ps[i] > maxDiff {
				maxDiff = ps[i+2] - ps[i]
			}
		}
		result := 0
		fmt.Println(maxDiff)

		if maxDiff == 0 && len(ps) == 2 {
			maxDiff = ps[1] - ps[0]
		}

		if maxDiff % 2 != 0 {
			result = (maxDiff + 1) / 2
		} else {
			result = maxDiff / 2
		}

		fmt.Println(result)

		results = append(results, strconv.Itoa(result))
	}

	result := strings.Join(results, "\n")

	writeResultFile(pathName + outputFileName, []byte(result))
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