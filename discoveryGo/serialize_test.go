package discoveryGo

import (
	"time"
	"encoding/json"
	"log"
	"fmt"
)

func Example_marshalJSON() {
	t := Task {
		"Laundry",
		DONE,
		NewDeadLine(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
		0,
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))

	// output:
	// {"Title":"Laundry","Status":2,"Deadline":"2015-08-16T15:43:00Z"}
}

func Example_unmarshalJSON() {
	b := []byte(`{"Title":"Laundry","Status":2,"Deadline":"2015-08-16T15:43:00Z"}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
	// output:
	// Laundry
	// 2
	// 2015-08-16 15:43:00 +0000 UTC
}

type MyStruct struct {
	Title		string	`json:"title"`
	Internal 	string	`json:"-"`
	Value 		int64	`json:",omitempty"`
	ID 			int64 	`json:",string"`
}

func Example_JSON_field() {
	myStruct := MyStruct{
		"My Title",
		"Internal strings....",
		12345,
		987654321,
	}
	b, err := json.Marshal(myStruct)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// output:
	// {"title":"My Title","Value":12345,"ID":"987654321"}
}

func Example_mapMarshalJSON() {
	b, _ := json.Marshal(map[string]string{
		"Name": "Johon",
		"Age": "16",
	})
	fmt.Println(string(b))

	c, _ := json.Marshal(map[string]interface{}{
		"Name": "Johon",
		"Age": "16",
	})
	fmt.Println(string(c))

	// output:
	// {"Age":"16","Name":"Johon"}
	// {"Age":"16","Name":"Johon"}
}

type Fields struct {
	VisibleField string `json:"visibleField"`
	InvisibleField string `json:"invisibleField"`
}

func ExampleOmitFields() {
	f := &Fields{"visible Text", "invisibleText"}
	b, _ := json.Marshal(
		struct {
			*Fields
			InvisibleField string `json:"invisibleField,omitempty"`
			Additional string `json:"additional,omitempty"`
		}{Fields: f, Additional: "additional Text"})
	fmt.Println(string(b))
	// output:
	// {"visibleField":"visible Text","additional":"additional Text"}
}