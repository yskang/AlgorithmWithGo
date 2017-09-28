package discoveryGo

import "fmt"

func (t Task2) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

func ExampleTask_String() {
	fmt.Println(Task{"Laundry", DONE, nil})
	// output:
	// [v] Laundry <nil>
}

type Task2 struct {
	Title		string	`json:"title,omitempty"`
	Status 		status 	`json:"status,omitempty"`
	Deadline 	*Deadline 	`json:"deadline,omitempty"`
	Priority 	int		`json:"priority,omitempty"`
	SubTasks 	[]Task2	`json:"subTasks,omitempty"`
}

type IncludeSubTasks Task2

func (t IncludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task2(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(prefix + "  ")
	}
	return str
}

func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}

func ExampleIncludeSubTasks_String() {
	fmt.Println(IncludeSubTasks(Task2{
		Title: "Laundry",
		Status: TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task2{{
			Title: "Wash",
			Status: TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []Task2{
				{"Put", DONE, nil, 2, nil},
				{"Detergent", TODO, nil, 2, nil},
			},
		},{
			Title: "Dry",
			Status: TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}, {
			Title: "Fold",
			Status: TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}))
	// output:
	// [ ] Laundry <nil>
	//   [ ] Wash <nil>
	//     [v] Put <nil>
	//     [ ] Detergent <nil>
	//   [ ] Dry <nil>
	//   [ ] Fold <nil>
}