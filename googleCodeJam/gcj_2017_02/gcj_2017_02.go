package gcj_2017_02

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func Gcj_2017_02 () {
	args := os.Args[1:]
	inputFileName := args[0]

	inputFileName = os.Getenv("GOPATH") +"/src/AlgorithmWithGo/gcj_2017_02/" + "B-large-practice.in"

	outputFileName := strings.Replace(inputFileName, ".in", ".out", 1)

	results := make([]string, 0)

	inputs := readFile(inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0 ; t < T ; t++ {
		k := inputs.Pop()
		tidyNumber := getTidyNumberLessThen(k)
		results = append(results, fmt.Sprintf("Case #%d: %s", t+1, tidyNumber))
	}

	result := strings.Join(results, "\n")

	writeResultFile(outputFileName, []byte(result))
}

func getTidyNumberLessThen(k string) string {
	ks := strings.Split(k, "")
	prev := ks[0]

	for i := 1 ; i < len(ks) ; i++ {
		if prev > ks[i] {
			ks[i-1] = string(prev[0] - 1)

			if ks[i-1] == "0" {
				ks[0] = ""
				for l := 1 ; l <= i-1 ; l++ {
					ks[l] = "9"
				}
			} else if i-2 >= 0 && ks[i-2] > ks[i-1] {
				temp := ks[i-1]
				l := i-2
				for l >= 0 && ks[l] > ks[l+1] {
					ks[l], ks[l+1] = temp, "9"
					l -= 1
				}

			}

			for j := i ; j < len(ks) ; j++ {
				ks[j] = "9"
			}
			return strings.Join(ks, "")
		}
		prev = ks[i]
	}

	return strings.Join(ks, "")
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