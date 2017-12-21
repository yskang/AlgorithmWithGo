package leetcode

import (
	"expvar"
	"strconv"
	"strings"
)

func DecodeString(s string) string {
	return decodeString(s)
}

const (
	STR = iota
	NUM = iota
	PAL = iota
	NON = iota
)

func decodeString(s string) string {
	stack := newStringStack()
	ans := make([]string, 0)

	for _, c := range s {
		if !stack.isEmpty() {
			prev := stack.pop()
			strings.
		} else {
			stack.push(string(c))
		}
	}

	return ans
}

type StringStask struct {
	strings []string
}

func newStringStack() StringStask {
	return StringStask{make([]string, 0)}
}

func (s *StringStask) push(str string) {
	s.strings = append(s.strings, str)
}

func (s *StringStask) pop() string {
	ans := s.strings[len(s.strings)-1]
	s.strings = append([]string{}, s.strings[:len(s.strings)-1]...)
	return ans
}

func (s *StringStask) isEmpty() bool {
	if len(s.strings) == 0 {
		return true
	}
	return false
}