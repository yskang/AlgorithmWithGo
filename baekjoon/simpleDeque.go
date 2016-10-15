// https://www.acmicpc.net/problem/10866

package baekjoon

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Deque struct {
	deque []int
}

func (d *Deque) pushFront(elem int) {
	d.deque = append([]int{elem}, d.deque...)
}

func (d *Deque) pushBack(elem int) {
	d.deque = append(d.deque, elem)
}

func (d *Deque) popFront() int {
	if len(d.deque) == 0 {
		return -1
	}

	front := d.deque[0]
	d.deque = d.deque[1:]
	return  front
}

func (d *Deque) popBack() int {
	if len(d.deque) == 0 {
		return -1
	}

	back := d.deque[len(d.deque) - 1]
	d.deque = d.deque[:len(d.deque) - 1]
	return back
}

func (d *Deque) size() int {
	return len(d.deque)
}

func (d *Deque) empty() int {
	if d.size() == 0 {
		return 1
	}
	return 0
}

func (d *Deque) front() int {
	if len(d.deque) == 0 {
		return -1
	}
	return d.deque[0]
}

func (d *Deque) back() int {
	if len(d.deque) == 0 {
		return -1
	}
	return d.deque[len(d.deque) - 1]
}

func SimpleDeque() {
	dq := Deque{make([]int, 0)}
	scanner := bufio.NewScanner(os.Stdin)

	nTest := 0
	fmt.Scanf("%d\n", &nTest)

	for scanner.Scan() {
		commStr := scanner.Text()
		commands := strings.Split(commStr, " ")

		switch commands[0] {
		case "push_back":
			val, _ := strconv.Atoi(commands[1])
			dq.pushBack(val)
		case "push_front":
			val, _ := strconv.Atoi(commands[1])
			dq.pushFront(val)
		case "front":
			fmt.Println(dq.front())
		case "back":
			fmt.Println(dq.back())
		case "size":
			fmt.Println(dq.size())
		case "empty":
			fmt.Println(dq.empty())
		case "pop_front":
			fmt.Println(dq.popFront())
		case "pop_back":
			fmt.Println(dq.popBack())
		}
	}
}
