package discoveryGo

import (
	"time"
	"fmt"
)

func ExampleDeadline_OverDue() {
	d1 := NewDeadLine(time.Now().Add(-4 * time.Hour))
	d2 := NewDeadLine(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, d1}
	t2 := Task{"4h later", TODO, d2}
	fmt.Println(d1.OverDue())
	fmt.Println(d2.OverDue())
	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	// Output:
	// true
	// false
	// true
	// false
}
