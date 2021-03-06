package discoveryGo

import (
	"strings"
	"sort"
	"fmt"
	"container/heap"
)

type CaseInsensitive []string

func (c CaseInsensitive) Len() int {
	return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) ||
		((strings.ToLower(c[i])) == strings.ToLower(c[j]) && c[i] < c[j])
}

func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func ExampleCaseInsensitive_sort() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStor",
	})

	sort.Sort(apple)
	fmt.Println(apple)
	// output:
	// [AppStor iPad iPhone MacBook]
}

func (c *CaseInsensitive) Push(x interface{}) {
	*c = append(*c, x.(string))
}

func (c *CaseInsensitive) Pop() interface{} {
	len := c.Len()
	last := (*c)[len-1]
	*c = (*c)[:len-1]
	return last
}

func ExampleCaseInsensitive_heap() {
	apple := CaseInsensitive([]string{"iPhone", "iPad", "MacBook", "AppStore"})
	heap.Init(&apple)
	for apple.Len() > 0 {
		poped := heap.Pop(&apple)
		s := poped.(string)
		fmt.Println(s)
	}
	// output:
	// AppStore
	// iPad
	// iPhone
	// MacBook
}

