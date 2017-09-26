package gcj_2017_01

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func Gcj_2017_01 () {
	args := os.Args[1:]
	inputFileName := args[0]

	inputFileName = os.Getenv("GOPATH") +"/src/AlgorithmWithGo/gcj_2017_01/" + "A-large-practice.in"

	outputFileName := strings.Replace(inputFileName, ".in", ".out", 1)

	results := make([]string, 0)

	inputs := readFile(inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0 ; t < T ; t++ {
		line := strings.Split(inputs.Pop(), " ")
		pancakeRow := line[0]
		k, _ := strconv.Atoi(line[1])

		count := flipPancakes(strings.Split(pancakeRow, ""), k)

		results = append(results, fmt.Sprintf("Case #%d: %s", t+1, count))
	}

	result := strings.Join(results, "\n")

	writeResultFile(outputFileName, []byte(result))
}
func flipPancakes(pancakeRow []string, k int) string {
	count := 0
	start := 0

	for  {
		if pancakeRow[start] == "-" {
			if start + k <= len(pancakeRow) {
				for i := start ; i < start + k ; i++ {
					if pancakeRow[i] == "-" {
						pancakeRow[i] = "+"
					} else {
						pancakeRow[i] = "-"
					}
				}
				count += 1
			} else {
				break
			}
		} else {
			start += 1
			if len(pancakeRow) - start < k {
				break
			}
		}
	}

	if strings.Contains(strings.Join(pancakeRow, ""), "-") {
		return "IMPOSSIBLE"
	}

	return strconv.Itoa(count)
}

func getScore(number string) {
	fmt.Println(number)
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