package discoveryGo

import (
	"errors"
	"fmt"
	"encoding/json"
	"net/http"
	"log"
)

type ID string

type DataAccess interface {
	Get(id ID) (Task, error)
	Put(id ID, t Task) error
	Post(t Task) (ID, error)
	Delete(id ID) error
}

type MemoryDataAccess struct {
	tasks map[ID]Task
	nextID int64
}

func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks: map[ID]Task{},
		nextID: int64(1),
	}
}

var ErrTaskNotExist = errors.New("task does not exist")

func (m *MemoryDataAccess) Get(id ID) (Task, error) {
	t, exist := m.tasks[id]
	if !exist {
		return Task{}, ErrTaskNotExist
	}
	return t, nil
}

func (m *MemoryDataAccess) Put(id ID, task Task) error {
	if _, exist := m.tasks[id]; !exist {
		return ErrTaskNotExist
	}
	m.tasks[id] = task
	return nil
}

func (m *MemoryDataAccess) Post(t Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID += 1
	m.tasks[id] = t
	return id, nil
}

func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exist := m.tasks[id]; !exist {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}

/*
		GET /api/v1/task/{id}

		PUT /api/v1/task/{id}
		task: {task}

		POST /api/v1/task/
		task: {task}

		DELETE /api/v1/task/{id}
 */

type ResponseError struct {
 	Err error
}

func (err ResponseError) MarshalJSON() ([]byte, error) {
 	if err.Err == nil {
 		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", err.Err)), nil
}

func (err *ResponseError) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	if v == nil {
		err.Err = nil
		return nil
	}
	switch tv := v.(type) {
	case string:
		if tv == ErrTaskNotExist.Error() {
			err.Err = ErrTaskNotExist
			return nil
		}
		err.Err = errors.New(tv)
		return nil
	default:
		return errors.New("ResponseError unmarshal failed")
	}
}

type Response struct {
	ID ID `json:"id,omitempty"`
	Task Task `json:"task"`
	Error ResponseError `json:"error"`
}


var m = NewMemoryDataAccess()
const pathPrefix = "/api/v1/task/"

func apiHandler(w http.ResponseWriter, r *http.Request) {
	getID := func() (ID, error) {
		id := ID(r.URL.Path[len(pathPrefix):])
		if id == "" {
			return id, errors.New("apiHandler: ID is empty")
		}
		return id, nil
	}

	getTask := func() ([]Task, error) {
		var result []Task
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
		encodedTasks, ok := r.PostForm["task"]
		if !ok {
			return nil, errors.New("task parameter expected")
		}
		for _, encodedTask := range encodedTasks {
			var t Task
			if err := json.Unmarshal([]byte(encodedTask), &t); err != nil {
				return nil, err
			}
			result = append(result, t)
		}
		return result, nil
	}

	switch r.Method {
	case "GET":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		t, err := m.Get(id)
		err = json.NewEncoder(w).Encode(Response{
			ID: id,
			Task: t,
			Error: ResponseError{err},
		})
		if err != nil {
			log.Println(err)
		}
	case "PUT":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		tasks, err := getTask()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			err = m.Put(id, t)
			err = json.NewEncoder(w).Encode(Response{
				ID: id,
				Task: t,
				Error: ResponseError{err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "POST":
		tasks, err := getTask()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			id, err := m.Post(t)
			err = json.NewEncoder(w).Encode(Response{
				ID: id,
				Task:t,
				Error:ResponseError{err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "DELETE":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		err = m.Delete(id)
		err = json.NewEncoder(w).Encode(Response{
			ID:id,
			Error:ResponseError{err},
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func taskMan() {
	http.HandleFunc(pathPrefix, apiHandler)
	log.Fatal(http.ListenAndServe(":8887", nil))
}

func ExampleTaskMan() {
	taskMan()
	// output:
	//
}
