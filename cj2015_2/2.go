package cj2015_2

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"sort"
)

func Cj2015_2 () {
	pathName := "/home/yskang/Downloads/2014_02/"
	inputFileName := "problem2.in"
	outputFileName := "result2.out"
	results := make([]string, 0)

	inputs := readFile(pathName + inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0 ; t < T ; t++ {
		result := "NO"
		firstLine := strings.Split(inputs.Pop(), " ")
		N, err := strconv.Atoi(firstLine[0])
		checkErr(err)
		K, err := strconv.Atoi(firstLine[1])
		checkErr(err)
		W, err := strconv.Atoi(firstLine[2])
		checkErr(err)

		traces := make([]Trace, 0)
		for i := 0 ; i < N ; i++ {
			traces = append(traces, newTrace(inputs.Pop()))
		}

		sort.Slice(traces, func(i, j int) bool {return traces[i].y < traces[j].y})

		for i := 0 ; i <= K ; i++ {
			if isValid(traces[i:len(traces)-K+i], W) {
				result = "YES"
				break
			}
		}
		results = append(results, result)
	}

	result := strings.Join(results, "\n")

	writeResultFile(pathName + outputFileName, []byte(result))
}

func isValid(traces []Trace, wallWidth int) bool {
	//fmt.Println("check : ", traces)
	outerTop := -9999999999
	outerBottom := 9999999999
	outerLeft := 9999999999
	outerRight := -9999999999

	for _, trace := range traces {
		if trace.y > outerTop {
			outerTop = trace.y
		}
		if trace.y < outerBottom {
			outerBottom = trace.y
		}
		if trace.x > outerRight {
			outerRight = trace.x
		}
		if trace.x < outerLeft {
			outerLeft = trace.x
		}
	}
	innerTop := outerTop - wallWidth
	innerBottm := outerBottom + wallWidth
	innerLeft := outerLeft + wallWidth
	innerRight := outerRight - wallWidth

	//fmt.Println("outter top, bottom, left, right", outerTop, outerBottom, outerLeft, outerRight)
	//fmt.Println("inner top, bottom, left, right", innerTop, innerBottm, innerLeft, innerRight)

	for _, trace := range traces {
		//fmt.Println("test trace ", trace)
		if  innerLeft < trace.x && trace.x < innerRight && innerBottm < trace.y && trace.y < innerTop {
			//fmt.Println("invalid!")
			return false
		}
	}
	//fmt.Println("all trace is valid!")
	return true
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

type Trace struct {
	x int
	y int
}

func newTrace(coordinate string) Trace {
	temp := strings.Split(coordinate, " ")
	x, _ := strconv.Atoi(temp[0])
	y, _ := strconv.Atoi(temp[1])
	return Trace{x, y}
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