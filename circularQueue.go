// https://www.acmicpc.net/problem/1021

package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

type CQueue struct {
	cQueue []int
}

func (c *CQueue) putBack(elem int) {
	c.cQueue = append(c.cQueue, elem)
}

func (c *CQueue) popFront() int {
	front := c.cQueue[0]
	c.cQueue = c.cQueue[1:]
	return front
}

func (c *CQueue) checkFront() int {
	return c.cQueue[0]
}

func (c *CQueue) checkLeftDistance(elem int) (int, bool) {
	for i := 0 ; i < len(c.cQueue) ; i++ {
		if elem == c.cQueue[i] {
			return i, true
		}
	}
	return 0, false
}

func (c *CQueue) checkRightDistance(elem int) (int, bool) {
	for i := 0 ; i < len(c.cQueue) ; i++ {
		if elem == c.cQueue[i] {
			return len(c.cQueue) - i, true
		}
	}
	return 0, false
}

func (c *CQueue) circulateLeft() {
	temp := c.cQueue[0]
	c.cQueue = c.cQueue[1:]
	c.cQueue = append(c.cQueue, temp)
}

func (c *CQueue) circulateRight() {
	temp := c.cQueue[len(c.cQueue) - 1]
	c.cQueue = append([]int{temp}, c.cQueue[:len(c.cQueue) - 1]...)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var sizeQ, numOut int
	qs := make([]int, 0)

	for scanner.Scan() {
		strIn := scanner.Text()
		ins := strings.Split(strIn, " ")
		sizeQ, _ = strconv.Atoi(ins[0])
		numOut, _ = strconv.Atoi(ins[1])
		break
	}

	for scanner.Scan() {
		strIn := scanner.Text()
		ins := strings.Split(strIn, " ")
		for _, v := range ins {
			iVal, _ := strconv.Atoi(v)
			qs = append(qs,iVal)
		}
		break
	}

	cQ := CQueue{make([]int, 0)}

	for i := 0 ; i < sizeQ ; i++ {
		cQ.putBack(i+1)
	}

	sum := 0
	for j := 0 ; j < numOut ; j++ {
		distLeft, _ := cQ.checkLeftDistance(qs[j])
		distRight, _ := cQ.checkRightDistance(qs[j])

		if distLeft < distRight {
			sum = sum + distLeft
			for k := 0 ; k < distLeft ; k++ {
				cQ.circulateLeft()
			}
		} else {
			sum = sum + distRight
			for k := 0 ; k < distRight ; k++ {
				cQ.circulateRight()
			}
		}
		cQ.popFront()

	}
	fmt.Println(sum)
}
