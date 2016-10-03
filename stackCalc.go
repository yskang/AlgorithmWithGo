package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

var priority = map[string]int{
	"(": -1,
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
}

type myStack struct {
	stack []string
}

func (s *myStack) Push(v string) {
	s.stack = append(s.stack, v)
}

func (s *myStack) Pop() string {

	if s.IsEmpty() == true {
		fmt.Printf("ROCK")
		os.Exit(0)
	}

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

func main() {
	var inStr string
	fmt.Scan(&inStr)
	//inStr := string("5+(1+2)*3")

	// make input equation to list
	tempStr := inStr

	tempStr = strings.Replace(tempStr, "(", ",(,", -1)
	tempStr = strings.Replace(tempStr, ")", ",),", -1)
	tempStr = strings.Replace(tempStr, "+", ",+,", -1)
	tempStr = strings.Replace(tempStr, "-", ",-,", -1)
	tempStr = strings.Replace(tempStr, "/", ",/,", -1)
	tempStr = strings.Replace(tempStr, "*", ",*,", -1)

	inBuffTemp := strings.Split(tempStr, ",")
	inBuff := make([]string, len(inBuffTemp))

	i := 0

	for _, str := range inBuffTemp {
		if str != "" {
			inBuff[i] = str
			i++
		}
	}

	// convert to postfix form
	stack := myStack{make([]string, 0)}

	outQueue := make([]string, 0)

	for _, tok := range inBuff {
		if tok == "" {
			break
		}
		if tok[0] >= '0' && tok[0] <= '9' {
			// if token is value, push to outQueue
			outQueue = append(outQueue, tok)
		} else if tok == "(" {
			stack.Push(tok)

		} else if tok == "+" || tok == "-" || tok == "*" || tok == "/" {
			for {
				if stack.IsEmpty() == true {
					stack.Push(tok)
					break
				} else {
					poped := stack.Pop()
					if checkPriority(poped, tok) == true {
						outQueue = append(outQueue, poped)
					} else {
						stack.Push(poped)
						stack.Push(tok)
						break
					}
				}
			}
		} else if tok == ")" {
			for {
				poped := stack.Pop()
				if poped == "(" {
					break
				} else {
					outQueue = append(outQueue, poped)
				}
			}
		}
	}

	for {
		if stack.IsEmpty() == true {
			break
		}
		outQueue = append(outQueue, stack.Pop())
	}

	// check postfix form
	//fmt.Println(outQueue)

	// calculate!!
	stack.Clear()

	for _, tok := range outQueue {
		if tok[0] >= '0' && tok[0] <= '9' {
			stack.Push(tok)
		} else {
			a := stack.Pop()
			b := stack.Pop()

			switch tok {
			case "+":
				stack.Push(BigAdd(a, b))
			case "-":
				stack.Push(BigSub(b, a))
			case "*":
				stack.Push(BigMul(a, b))
			case "/":
				if a == 0 {
					fmt.Printf("ROCK")
					return
				}
				stack.Push(BigDiv(b, a))
			}
		}
	}

	fmt.Printf(stack.Pop())
}
func BigDiv(a string, b string) string {
	return "1"
}
func BigMul(a string, b string) string {
	valA, signA := GetBigInt(a)
	valB, signB := GetBigInt(b)




	resultString := ""

	if signA != signB {
		return "-" + resultString
	}

	return resultString
}

func BigSub(a string, b string) string {
	valA, signA := GetBigInt(a)
	valB, signB := GetBigInt(b)

	if signB == false {
		return BigAdd(a, valB)
	} else if signA == false && signB == true {
		return "-" + BigAdd(valA, valB)
	}

	if CompareAandB(valA, valB) == "<" {
		return "-" + BigSub(valB, valA)
	} else if CompareAandB(valA, valB) == "=" {
		return "0"
	}

	results := make([]byte, 0)

	carry := false
	for i := 0 ; true ; i++ {
		var aa, bb byte

		if i < len(valA) {
			aa = valA[len(valA)-i-1]
		} else {
			aa = '0'
		}

		if i < len(valB) {
			bb = valB[len(valB)-i-1]
		} else {
			bb = '0'
		}

		if carry == true {
			aa = aa - 1
		}

		if aa < bb {
			results = append(results, aa + byte(10) - bb)
			carry = true
		} else {
			results = append(results, aa - bb)
			carry = false
		}

		if len(valA) == i {
			if results[i-1] == 0 {
				results = results[:i-1]
			} else if results[i] == 0 {
				results = results[:i]
			}
			break
		}
	}

	resultString := ""

	for j := len(results)-1 ; j >= 0 ; j-- {
		resultString = resultString + strconv.Itoa(int(results[j]))
	}

	return resultString
}

func CompareAandB(a string, b string) string {
	if len(a) < len(b) {
		return "<"
	} else if len(a) > len(b) {
		return ">"
	}

	for i := 0 ; i < len(a) ; i++ {
		if a[i] > b[i] {
			return ">"
		} else if a[i] < b[i] {
			return "<"
		}
	}

	return "="
}

func GetBigInt(intString string) (string, bool) {
	if intString[0] == '-' {
		return intString[1:], false
	}
	return intString, true
}

func BigAdd(a string, b string) string {
	valA, signA := GetBigInt(a)
	valB, signB := GetBigInt(b)

	if signA == false && signB == false {
		return "-" + BigAdd(valA, valB)
	} else if signA == false && signB == true {
		return BigSub(valB, valA)
	} else if signA == true && signB == false {
		return BigSub(valA, valB)
	}

	results := make([]uint8, 0)
	carry := 0

	lenA := len(a)
	lenB := len(b)
	lenBig := 0

	if lenA > lenB {
		lenBig = lenA
	} else {
		lenBig = lenB
	}


	for i := 0 ; true ; i++ {
		var aa, bb byte

		if i < lenA {
			aa = a[lenA-i-1]
		} else {
			aa = '0'
		}

		if i < lenB {
			bb = b[lenB-i-1]
		} else {
			bb = '0'
		}

		sum := int(aa) - '0' + int(bb) - '0' + carry
		carry = sum / 10
		sum = sum % 10
		results = append(results, byte(sum))

		if lenBig == i {
			if results[i] == 0 {
				results = results[:i]
			}
			break
		}
	}

	resultString := ""

	for j := len(results)-1 ; j >= 0 ; j-- {
		resultString = resultString + strconv.Itoa(int(results[j]))
	}

	return resultString
}

func checkPriority(poped string, tok string) bool{
	return priority[poped] >= priority[tok]
}

