package myLibs

type myStack struct {
	stack []string
}

func NewStack() *myStack {
	return &myStack{make([]string, 0)}
}

func (s *myStack) Push(v string) {
	s.stack = append(s.stack, v)
}

func (s *myStack) Pop() string {
	topVal := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return topVal
}

func (s *myStack) IsEmpty() bool {
	if len(s.stack) > 0 {
		return  false
	}
	return true
}

func (s *myStack) Clear() {
	for {
		if s.IsEmpty() == true {
			return
		}
		s.Pop()
	}
}
