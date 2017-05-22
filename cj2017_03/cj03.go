package cj2017_03

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func CJ_2017_03() {
	args := os.Args[1:]
	inputFileName := args[0]
	outputFileName := strings.Replace(inputFileName, ".in", ".out", 1)

	results := make([]string, 0)

	inputs := readFile(inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0 ; t < T ; t++ {
		line := strings.Split(inputs.Pop(), " ")
		n, err := strconv.Atoi(line[0])
		checkErr(err)
		k, err := strconv.Atoi(line[1])
		checkErr(err)

		fmt.Println(n, k)

		results = append(results, isValid(n, k))
	}

	result := strings.Join(results, "\n")

	writeResultFile(outputFileName, []byte(result))
}

func isValid(n int, k int) string {
	if n == k {
		return "O"
	}

	var remainder int

	ten := n / 10

	for i := ten; i >= 0 ; i-- {
		remainder = n - i * 10
		nine := remainder / 9

		if i * 2 <= k {
			for j := nine ; j >= 0 ; j-- {
				r := remainder - j * 9
				five := r / 5

				if i * 2 + j * 3 <= k {
					for l := five ; l >= 0 ; l-- {
						s := r - l * 5
						four := s / 4

						if i * 2 + j * 3 + l * 2 <= k {
							for m := four ; m >= 0 ; m-- {
								t := s - m * 4

								if i * 2 + j * 3 + l * 2 + m * 3 + t == k {
									return "O"
								}
							}
						}
					}
				}

			}
		}

	}

	return "X"
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