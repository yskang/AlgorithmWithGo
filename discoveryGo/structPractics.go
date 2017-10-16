package discoveryGo

import (
	"time"
	"errors"
	"fmt"
)

type Deadline struct {
	time.Time
}

func NewDeadLine(t time.Time) *Deadline {
	return &Deadline{t}
}

func (d Deadline) OverDue() bool {
	return d.Before(time.Now())
}

type status int

const (
	UNKNOWN	status = iota
	TODO
	DONE
)

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int		`json:"priority,omitempty"`
}

func (s status) String() string {
	switch s {
	case UNKNOWN:
		return "UNKNOWN"
	case TODO:
		return "TODO"
	case DONE:
		return "DONE"
	default:
		return ""
	}
}

func (s status) MarshalJSON() ([]byte, error) {
	str := s.String()
	if str == "" {
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
	return []byte(fmt.Sprintf("\"%s\"", str)), nil
}

func (t Task) OverDue() bool {
	return t.Deadline.OverDue()
}