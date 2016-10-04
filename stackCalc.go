package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

	topVal := s.stack[len(s.stack) - 1]
	s.stack = s.stack[:len(s.stack) - 1]
	return topVal
}

func (s *myStack) IsEmpty() bool {
	if len(s.stack) > 0 {
		return false
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
				if a == "0" {
					fmt.Printf("ROCK")
					return
				}
				val, _ := BigDiv(b, a)
				stack.Push(val)
			}
		}
	}

	fmt.Printf(stack.Pop())
}

func BigDiv(a string, b string) (string, string){
	valA, signA := GetBigInt(a)
	valB, signB := GetBigInt(b)

	lenA := len(valA)
	lenB := len(valB)

	if signA != signB {
		val, remain := BigDiv(valA, valB)
		return "-" + val, remain
	}

	switch CompareAandB(valA, valB) {
		case "<":
			return "0", valA
		case "=":
			return "1", "0"
	}

	resultStr := ""

	for i := 0 ; i < lenA - lenB + 1 ; i++ {
		for j := 0 ; j <= 10 ; j++ {
			if CompareAandB(BigMul(valB, strconv.Itoa(j)), valA[0:i + lenB]) == ">" {
				resultStr = resultStr + strconv.Itoa(j-1)
				temp := valB
				for k := 0 ; k < lenA - lenB - i ; k++ {
					temp = temp + "0"
				}
				valA = BigSub(valA, BigMul(strconv.Itoa(j-1), temp))
				valA = padZero(valA, lenA)
				break
			}
		}
	}

	return removeFrontZero(resultStr), removeFrontZero(valA)
}
func padZero(val string, length int) string {
	for {
		if len(val) == length {
			return val
		}
		val = "0" + val
	}
}

func BigMul(a string, b string) string {
	valA, signA := GetBigInt(a)
	valB, signB := GetBigInt(b)

	resultString := "0"
	tempResults := make([]string, 0)

	if signA != signB {
		return "-" + BigMul(valA, valB)
	}

	for i := len(valB) - 1 ; i >= 0 ; i-- {
		for j := len(valA) - 1 ; j >= 0 ; j-- {
			temp := strconv.Itoa(int(valA[j] - '0') * int(valB[i] - '0'))
			for k := 0 ; k < len(valA) - j - 1 + len(valB) - i - 1; k++ {
				temp = temp + "0"
			}
			tempResults = append(tempResults, temp)
		}
	}

	for _, tempResult := range tempResults {
		resultString = BigAdd(resultString, tempResult)
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
	for i := 0; true; i++ {
		var aa, bb byte

		if i < len(valA) {
			aa = valA[len(valA) - i - 1]
		} else {
			aa = '0'
		}

		if i < len(valB) {
			bb = valB[len(valB) - i - 1]
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
			break
		}
	}

	resultString := ""

	for j := len(results) - 1; j >= 0; j-- {
		resultString = resultString + strconv.Itoa(int(results[j]))
	}

	return removeFrontZero(resultString)
}

func CompareAandB(a string, b string) string {
	valA, _ := GetBigInt(a)
	valB, _ := GetBigInt(b)

	if len(valA) < len(valB) {
		return "<"
	} else if len(valA) > len(valB) {
		return ">"
	}

	for i := 0; i < len(valA); i++ {
		if valA[i] > valB[i] {
			return ">"
		} else if valA[i] < valB[i] {
			return "<"
		}
	}

	return "="
}

func GetBigInt(intString string) (string, bool) {
	if intString[0] == '-' {
		return removeFrontZero(intString[1:]), false
	}
	return removeFrontZero(intString), true
}

func removeFrontZero(intString string) string {
	for i, val := range intString {
		if val != '0' {
			return intString[i:]
		}
	}
	return "0"
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

	lenA := len(valA)
	lenB := len(valB)
	lenBig := 0

	if lenA > lenB {
		lenBig = lenA
	} else {
		lenBig = lenB
	}

	for i := 0; true; i++ {
		var aa, bb byte

		if i < lenA {
			aa = valA[lenA - i - 1]
		} else {
			aa = '0'
		}

		if i < lenB {
			bb = valB[lenB - i - 1]
		} else {
			bb = '0'
		}

		sum := int(aa) - '0' + int(bb) - '0' + carry
		carry = sum / 10
		sum = sum % 10
		results = append(results, byte(sum))

		if lenBig == i {
			break
		}
	}

	resultString := ""

	for j := len(results) - 1; j >= 0; j-- {
		resultString = resultString + strconv.Itoa(int(results[j]))
	}

	return removeFrontZero(resultString)
}

func checkPriority(poped string, tok string) bool {
	return priority[poped] >= priority[tok]
}
