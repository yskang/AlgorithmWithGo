// https://www.acmicpc.net/problem/1966

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Doc struct {
	name     int
	priority int
}

type MyQueue struct{
	myQueue []Doc
}

func (q *MyQueue) push(data Doc) {
	q.myQueue = append(q.myQueue, data)
}

func (q *MyQueue) pop() (Doc, bool) {
	if len(q.myQueue) == 0 {
		return Doc{0, 0}, false
	}

	result := q.myQueue[0]
	q.myQueue = q.myQueue[1:]
	return result, true
}

func (q *MyQueue) size() int {
	return len(q.myQueue)
}

func (q *MyQueue) isEmpty() bool {
	if len(q.myQueue) == 0 {
		return true
	}
	return false
}

func (q *MyQueue) front() (Doc, bool) {
	if len(q.myQueue) == 0 {
		return Doc{0, 0}, false
	}
	return q.myQueue[0], true
}

func (q *MyQueue) back() (Doc, bool) {
	if len(q.myQueue) == 0 {
		return Doc{0, 0}, false
	}
	return q.myQueue[len(q.myQueue) - 1], true
}

func isBigPriority(p int, docs []Doc) bool {
	for _, doc := range docs {
		if doc.priority > p {
			return true
		}
	}
	return false
}

func main() {
	numOfTest := 0
	fmt.Scanf("%d\n", &numOfTest)
	reader := bufio.NewReader(os.Stdin)
	var pStr string
	for i := 0 ; i < numOfTest ; i++ {
		printCount := 0
		var num, pos int
		fmt.Scanf("%d %d\n", &num, &pos)
		pStr, _ = reader.ReadString('\n')
		pStr = strings.Replace(pStr, "\n", "", -1)

		queue := MyQueue{make([]Doc, 0)}
		for i, v := range strings.Split(pStr, " ") {
			val, _ := strconv.Atoi(v)
			queue.push(Doc{i, val})
		}

		for {
			doc, _ := queue.pop()

			if isBigPriority(doc.priority, queue.myQueue) == true {
				queue.push(doc)
			} else {
				printCount++
				if doc.name == pos {
					fmt.Println(printCount)
					break
				}
			}

		}
	}

}
