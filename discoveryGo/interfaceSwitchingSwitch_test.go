package discoveryGo

import (
	"github.com/jaeyeom/gogo/task"
	"fmt"
	"strconv"
	"strings"
)

func ExampleJoin() {
	t := task.Task{
		Title: "Laundry",
		Status: task.DONE,
		Deadline: nil,
	}
	fmt.Println(Join(",", 1, "two", 3, t))
	// output:
	// 1,two,3,[v] Laundry <nil>
}

func Join(sep string, a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}
	t := make([]string, len(a))
	for i := range a {
		switch x := a[i].(type) {
		case string:
			t[i] = x
		case int:
			t[i] = strconv.Itoa(x)
		case fmt.Stringer:
			t[i] = x.String()
		}
	}
	return strings.Join(t, sep)
}
