package main

import (
	"fmt"
	"strings"
	"AlgorithmWithGo/myLibs"
)

var priority = map[string]int{
	"(": -1,
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
}

func main() {
	inStr := string("55+((12+23)*42)-3323")

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
	stack := myLibs.NewStack()

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

	outQueue = append(outQueue, stack.Pop())

	// calculate!!


	fmt.Println(outQueue)
}

func checkPriority(poped string, tok string) bool{
	return priority[poped] >= priority[tok]
}

