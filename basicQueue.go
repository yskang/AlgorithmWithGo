// https://www.acmicpc.net/problem/10845

package main

import (
	"fmt"
	"strings"
	"strconv"
)

type MyQueue struct{
	myQueue []int
}

func (q *MyQueue) push(data int) {
	q.myQueue = append(q.myQueue, data)
}

func (q *MyQueue) pop() int {
	if len(q.myQueue) == 0 {
		return -1
	}

	result := q.myQueue[0]
	q.myQueue = q.myQueue[1:]
	return result
}

func (q *MyQueue) size() int {
	return len(q.myQueue)
}

func (q *MyQueue) empty() int {
	if len(q.myQueue) == 0 {
		return 1
	}
	return 0
}

func (q *MyQueue) front() int {
	if len(q.myQueue) == 0 {
		return -1
	}
	return q.myQueue[0]
}

func (q *MyQueue) back() int {
	if len(q.myQueue) == 0 {
		return -1
	}
	return q.myQueue[len(q.myQueue) - 1]
}

func main() {
	outStr := make([]string, 0)
	queue := MyQueue{make([]int, 0)}
	inComm := ""
	inData := 0
	num := 0
	fmt.Scanf("%d\n", &num)
	for i := 0 ; i < num; i++ {
		fmt.Scanf("%s %d\n", &inComm, &inData)
		switch inComm {
		case "push":
			queue.pushPQ(inData)
		case "pop":
			outStr = append(outStr, strconv.Itoa(queue.popPQ()))
		case "size":
			outStr = append(outStr, strconv.Itoa(queue.sizePQ()))
		case "empty":
			outStr = append(outStr, strconv.Itoa(queue.empty()))
		case "front":
			outStr = append(outStr, strconv.Itoa(queue.frontPQ()))
		case "back":
			outStr = append(outStr, strconv.Itoa(queue.backPQ()))
		}
	}

	fmt.Println(strings.Join(outStr, "\n"))
}
