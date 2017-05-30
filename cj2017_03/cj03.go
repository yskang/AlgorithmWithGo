package cj2017_03

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"sync"
	"runtime"
)

//2m 32s 684ms (3m 4s 660ms)
func CJ_2017_03() {
	args := os.Args[1:]
	inputFileName := args[0]

	inputFileName = os.Getenv("GOPATH") +"/src/AlgorithmWithGo/cj2017_03/" + "Set3.in"
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

// 1m 28s 583ms (2m 31s 709ms)
func CJ_2017_03_with_goroutine() {
	runtime.GOMAXPROCS(4)

	args := os.Args[1:]
	inputFileName := args[0]

	inputFileName = os.Getenv("GOPATH") +"/src/AlgorithmWithGo/cj2017_03/" + "Set3.in"
	outputFileName := strings.Replace(inputFileName, ".in", "withGoroutine.out", 1)


	inputs := readFile(inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	var wait sync.WaitGroup
	wait.Add(T)

	results := make([]string, T)

	for t := 0 ; t < T ; t++ {
		line := strings.Split(inputs.Pop(), " ")
		n, err := strconv.Atoi(line[0])
		checkErr(err)
		k, err := strconv.Atoi(line[1])
		checkErr(err)

		fmt.Println(n, k)

		go isValidGoRoutine(n, k, t, results, wait.Done)
	}

	wait.Wait()
	result := strings.Join(results, "\n")

	writeResultFile(outputFileName, []byte(result))
}

func CJ_2017_03_with_goroutine_channel() {
	runtime.GOMAXPROCS(4)

	args := os.Args[1:]
	inputFileName := args[0]

	inputFileName = os.Getenv("GOPATH") +"/src/AlgorithmWithGo/cj2017_03/" + "Set3.in"
	outputFileName := strings.Replace(inputFileName, ".in", "withGoroutine.out", 1)


	inputs := readFile(inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	results := make([]string, T)

	channels := make([]chan string, T)

	for t := 0 ; t < T ; t++ {
		line := strings.Split(inputs.Pop(), " ")
		n, err := strconv.Atoi(line[0])
		checkErr(err)
		k, err := strconv.Atoi(line[1])
		checkErr(err)

		channels[t] = make(chan string, 1)

		go func(n, k, t int) {
			fmt.Println(n, k)

			time.Sleep(0)
			isValidGoRoutineChannel(n, k, channels[t])
		} (n, k, t)
		time.Sleep(0)
	}

	for i := 0 ; i < T ; i++ {
		results[i] = <- channels[i]
	}

	result := strings.Join(results, "\n")

	writeResultFile(outputFileName, []byte(result))
}

func isValidGoRoutineChannel(n int, k int, ch chan string) {
	if n == k {
		ch <- "O"
		return
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
									ch <- "O"
									return
								}
							}
						}
					}
				}

			}
		}

	}

	ch <- "X"
	return
}


func isValidGoRoutine(n int, k int, testNum int, results []string, done func()) {
	defer done()
	if n == k {
		results[testNum] = "O"
		return
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
									results[testNum] = "O"
									return
								}
							}
						}
					}
				}

			}
		}

	}

	results[testNum] = "X"
	return
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