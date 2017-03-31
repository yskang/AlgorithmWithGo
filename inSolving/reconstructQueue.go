package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reconstructQueue([][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}))
}

type People [][]int

func (slice People) Len() int {
	return len(slice)
}

func (slice People) Less(i, j int) bool {
	return slice[i][1] < slice[j][1]
}

func (slice People) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func reconstructQueue(people [][]int) [][]int {
	p := People{}
	p = people
	sort.Sort(p)
	fmt.Println(people)

	for i := 0 ; i < len(people) ; i++ {
		if !isSatisfied(people, i) {
			fmt.Println("not satisfied!!", people[i])
			moveToProperPosition(people, i)
			i = 0
		}
	}

	return people
}

func isSatisfied(people [][]int, index int) bool {
	count := 0

	for i := 0 ; i < index ; i++ {
		if people[index][0] <= people[i][0] {
			count += 1
		}
	}

	return count == people[index][1]
}

func moveToProperPosition(people [][]int, index int) {
	fmt.Println("move to proper position!", people[index])
	count := 0
	for i := 0 ; i < index ; i++ {
		if people[index][0] <= people[i][0] {
			count += 1
		}
		if people[index][1] == count {
			fmt.Println("move to ", i)
			temp := make([]int, 2)
			copy(temp, people[index])

			for j := index - 1 ; j >= i + 1 ; j-- {
				people[j+1] = people[j]
			}
			people[i+1] = temp
			return
		}
	}
}
