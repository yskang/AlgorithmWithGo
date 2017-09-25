package discoveryGo

import (
	"time"
)

type Deadline time.Time

func NewDeadLine(t time.Time) Deadline {
	return Deadline(t)
}

func (d Deadline) OverDue() bool {
	return time.Time(d).Before(time.Now())
}

type status int

const (
	UNKNOWN	status = iota
	TODO
	DONE
)

type Task struct {
	Title string
	Status status
	Deadline Deadline
}

func (t Task) OverDue() bool {
	return t.Deadline.OverDue()
}