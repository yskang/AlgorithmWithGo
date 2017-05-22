package cj2017_04

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func CJ_2017_04() {
	args := os.Args[1:]
	inputFileName := args[0]
	outputFileName := strings.Replace(inputFileName, ".in", ".out", 1)

	results := make([]string, 0)

	inputs := readFile(inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0 ; t < T ; t++ {
		n := make([]int, 5)
		numsString := strings.Split(inputs.Pop(), " ")
		for i, num := range numsString {
			n[i], _ = strconv.Atoi(num)
		}

		ith, repeatNumber := getRepeatNumber(n)

		results = append(results, ith + " " + repeatNumber)
	}

	result := strings.Join(results, "\n")

	writeResultFile(outputFileName, []byte(result))
}

func getRepeatNumber(numbers []int) (string, string) {
	DIV := 1000000007
	parts := make([]int, 4)
	parts[0] = numbers[0]
	parts[1] = numbers[1] + parts[0]
	parts[2] = numbers[2] + parts[1]
	parts[3] = numbers[3] + parts[2]
	oneSet := parts[3] + numbers[4]
	fmt.Println(numbers)
	fmt.Println(parts)

	baseRemainder := 11 % oneSet
	baseQuotient := 11 / oneSet
	remainder := baseRemainder
	quotient := baseQuotient
	fmt.Println(baseRemainder, baseQuotient)

	for i := 2 ; true ; i++ {
		for j := 0 ; j < 9 ; j++ {

			//fmt.Println("============", strings.Repeat(strconv.Itoa(j+1),i), "=============")
			//fmt.Println("one set", oneSet)
			//fmt.Println("base remainder", baseRemainder, "base quotient", baseQuotient)
			//fmt.Println("remainder", remainder, "quotient", quotient)
			//fmt.Println("===============================")

			if remainder != 0 {
				for k := 0 ; k < 4 ; k++ {
					if remainder == parts[k] {
						fmt.Println("found!! repeat number", (quotient * 5 + k) % DIV , j +1, "(", i, ")")
						return fmt.Sprintf("%d", (quotient * 5 + k) % DIV), fmt.Sprintf("%d(%d)", j+1, i)
					}
				}
			} else {
				fmt.Println("found!! repeat number", (quotient * 5 - 1) % DIV, j+1,  "(", i, ")")
				return fmt.Sprintf("%d", (quotient * 5 - 1) % DIV), fmt.Sprintf("%d(%d)", j+1, i)
			}

			remainder += baseRemainder
			quotient = (quotient % DIV + baseQuotient % DIV) % DIV
			if remainder >= oneSet {
				remainder = remainder - oneSet
				quotient = (quotient % DIV + 1) % DIV
			}
		}

		baseRemainder = (baseRemainder * 10) + 1
		baseQuotient = ((baseQuotient % DIV * 10) % DIV + baseRemainder / oneSet) % DIV
		baseRemainder = baseRemainder % oneSet

		remainder += 1
		if remainder >= oneSet {
			remainder = remainder - oneSet
			quotient += 1
		}
	}

	fmt.Println("fail to find")
	return "", ""
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