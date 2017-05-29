package gcj_2017_03

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func Gcj_2017_03 () {
	args := os.Args[1:]
	inputFileName := args[0]

	inputFileName = os.Getenv("GOPATH") +"/src/AlgorithmWithGo/gcj_2017_03/" + "C-small-practice-1.in"

	outputFileName := strings.Replace(inputFileName, ".in", ".out", 1)

	results := make([]string, 0)

	inputs := readFile(inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0 ; t < T ; t++ {
		line := strings.Split(inputs.Pop(), " ")
		n, _ := strconv.Atoi(line[0])
		k, _ := strconv.Atoi(line[1])
		max, min := calculateEmpty(n, k)
		results = append(results, fmt.Sprintf("Case #%d: %d %d", t+1, max, min))
	}

	result := strings.Join(results, "\n")

	writeResultFile(outputFileName, []byte(result))
}

func calculateEmpty(n int, k int) (int, int) {
	return 0, 0
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