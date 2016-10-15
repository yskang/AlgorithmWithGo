// https://www.acmicpc.net/problem/10845
package baekjoon

import (
	"fmt"
	"strings"
	"strconv"
)

type BQueue struct{
	myQueue []int
}

func (q *BQueue) push(data int) {
	q.myQueue = append(q.myQueue, data)
}

func (q *BQueue) pop() int {
	if len(q.myQueue) == 0 {
		return -1
	}

	result := q.myQueue[0]
	q.myQueue = q.myQueue[1:]
	return result
}

func (q *BQueue) size() int {
	return len(q.myQueue)
}

func (q *BQueue) empty() int {
	if len(q.myQueue) == 0 {
		return 1
	}
	return 0
}

func (q *BQueue) front() int {
	if len(q.myQueue) == 0 {
		return -1
	}
	return q.myQueue[0]
}

func (q *BQueue) back() int {
	if len(q.myQueue) == 0 {
		return -1
	}
	return q.myQueue[len(q.myQueue) - 1]
}

func BasicQueue() {
	outStr := make([]string, 0)
	queue := BQueue{make([]int, 0)}
	inComm := ""
	inData := 0
	num := 0
	fmt.Scanf("%d\n", &num)
	for i := 0 ; i < num; i++ {
		fmt.Scanf("%s %d\n", &inComm, &inData)
		switch inComm {
		case "push":
			queue.push(inData)
		case "pop":
			outStr = append(outStr, strconv.Itoa(queue.pop()))
		case "size":
			outStr = append(outStr, strconv.Itoa(queue.size()))
		case "empty":
			outStr = append(outStr, strconv.Itoa(queue.empty()))
		case "front":
			outStr = append(outStr, strconv.Itoa(queue.front()))
		case "back":
			outStr = append(outStr, strconv.Itoa(queue.back()))
		}
	}

	fmt.Println(strings.Join(outStr, "\n"))
}
