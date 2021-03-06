// https://www.acmicpc.net/problem/1966

package baekjoon

import (
	"fmt"
)

type Doc struct {
	name     int
	priority int
}

type PQueue struct{
	myQueue []Doc
}

func (q *PQueue) pushPQ(data Doc) {
	q.myQueue = append(q.myQueue, data)
}

func (q *PQueue) popPQ() (Doc, bool) {
	if len(q.myQueue) == 0 {
		return Doc{0, 0}, false
	}

	result := q.myQueue[0]
	q.myQueue = q.myQueue[1:]
	return result, true
}

func (q *PQueue) sizePQ() int {
	return len(q.myQueue)
}

func (q *PQueue) isEmpty() bool {
	if len(q.myQueue) == 0 {
		return true
	}
	return false
}

func (q *PQueue) frontPQ() (Doc, bool) {
	if len(q.myQueue) == 0 {
		return Doc{0, 0}, false
	}
	return q.myQueue[0], true
}

func (q *PQueue) backPQ() (Doc, bool) {
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

func PrinterQueue() {
	numOfTest := 0
	fmt.Scanf("%d\n", &numOfTest)
	for i := 0 ; i < numOfTest ; i++ {
		printCount := 0
		var num, pos int
		fmt.Scanf("%d %d\n", &num, &pos)
		queue := PQueue{make([]Doc, 0)}
		for j := 0 ; j < num ; j++ {
			var docPriority int
			fmt.Scanf("%d", &docPriority)
			queue.pushPQ(Doc{j, docPriority})
		}

		for {
			doc, _ := queue.popPQ()

			if isBigPriority(doc.priority, queue.myQueue) == true {
				queue.pushPQ(doc)
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
