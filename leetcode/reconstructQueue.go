package leetcode

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)


func ReconstructQueue() {
	testInput := "[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]"
	testData := make([][]int, 0)
	temp := strings.Split(testInput, ",")
	var pair []int
	for i, t := range temp {
		t = strings.Replace(t, "[", "", -1)
		t = strings.Replace(t, "]", "", -1)
		tInt, _ := strconv.Atoi(t)
		if i % 2 == 0 {
			pair = make([]int, 2)
			pair[0] = tInt
		} else {
			pair[1] = tInt
			testData = append(testData, pair)
		}
	}

	fmt.Println(reconstructQueue(testData))
}

type People [][]int

func (slice People) Len() int {
	return len(slice)
}

func (slice People) Less(i, j int) bool {
	if slice[i][0] == slice[j][0] {
		return slice[i][1] < slice[j][1]
	}
	return slice[i][0] > slice[j][0]
}

func (slice People) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func reconstructQueue(people [][]int) [][]int {
	p := People{}
	p = people
	sort.Sort(p)

	rearranged := make([][]int, 0)

	for _, person := range people {
		rearranged = append(rearranged[:person[1]], append([][]int{person} , rearranged[person[1]:]...)...)
	}

	return rearranged
}
