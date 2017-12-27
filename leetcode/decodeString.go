package leetcode

import (
	"strconv"
	"strings"
)

func DecodeString(s string) string {
	return decodeString(s)
}

const (
	STR = iota
	NUM = iota
	BR  = iota
)

func decodeString(s string) string {
	ans := make([]string, 0)
	stack := newStringStack()

	for _, c := range s {
		if stack.isEmpty() {
			stack.push(string(c))
		} else {
			if getKind(c) == BR {
				if c == '[' {
					stack.push(string(c))
				} else {
					temp := stack.pop()
					stack.pop()
					num, _ := strconv.Atoi(stack.pop())
					if stack.isEmpty() {
						ans = append(ans, strings.Repeat(temp, num))
					} else {
						tail := strings.Repeat(temp, num)
						if getKind(rune(stack.checkTop()[0])) == STR {
							head := stack.pop()
							stack.push(head + tail)
						} else {
							stack.push(tail)
						}
					}
				}
			} else if getKind(rune(stack.checkTop()[0])) == getKind(c) {
				stack.push(stack.pop() + string(c))
			} else {
				stack.push(string(c))
			}
		}
	}

	// fmt.Println(stack.String())
	if !stack.isEmpty() {
		return strings.Join(ans, "") + stack.pop()
	}
	return strings.Join(ans, "")
}

func getKind(c rune) int {
	if '0' <= c && c <= '9' {
		return NUM
	} else if 'a' <= c && c <= 'z' {
		return STR
	}
	return BR
}

type StringStack struct {
	strs []string
}

func newStringStack() StringStack {
	return StringStack{make([]string, 0)}
}

func (s *StringStack) push(str string) {
	s.strs = append(s.strs, str)
}

func (s *StringStack) pop() string {
	ans := s.strs[len(s.strs)-1]
	s.strs = append([]string{}, s.strs[:len(s.strs)-1]...)
	return ans
}

func (s *StringStack) checkTop() string {
	return s.strs[len(s.strs)-1]
}

func (s *StringStack) isEmpty() bool {
	if len(s.strs) == 0 {
		return true
	}
	return false
}

func (s *StringStack) String() string {
	return "B {" + strings.Join(s.strs, " | ") + "} T"
}
