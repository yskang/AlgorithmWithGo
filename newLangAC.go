package main

import (
	"fmt"
	"strconv"
	"strings"
)

func printArray(arr []int, isReverse bool) {
	arrStr := "["
	var index int
	for i := 0 ; i < len(arr) ; i++ {
		if isReverse {
			index = len(arr) - i - 1
		} else {
			index = i
		}
		arrStr = arrStr + strconv.Itoa(arr[index])
		if i != len(arr) - 1 {
			arrStr = arrStr + ","
		}
	}
	arrStr = arrStr + "]"
	println(arrStr)
}

func delFirst(arr []int, isReverse bool) ([]int, bool) {
	if len(arr) == 0 {
		return arr, false
	}

	if isReverse {
		return arr[:len(arr) - 1], true
	} else {
		return arr[1:], true
	}
}

func main() {
	numTest := 0
	fmt.Scanf("%d\n", &numTest)

	for i := 0 ; i < numTest ; i++ {
		valid := true
		isReverse := false
		command := ""
		arrSize := 0
		inStr := ""
		array := make([]int, 0)
		fmt.Scanf("%s\n", &command)
		fmt.Scanf("%d\n", &arrSize)
		fmt.Scanf("[%s]\n", &inStr)

		inStr = strings.Replace(inStr, "]", "", 1)
		inStrArray := strings.Split(inStr, ",")

		for _, v := range inStrArray {
			v = strings.TrimSpace(v)
			intVal, err := strconv.Atoi(v)
			if err == nil {
				array = append(array, intVal)
			} else {
				fmt.Println("error!!", v)
				valid = false
				break
			}
		}

		if len(array) != arrSize || valid == false{
			valid = false
		} else {
			for _, com := range command {
				switch com {
				case 'R':
					isReverse = !isReverse
				case 'D':
					array, valid = delFirst(array, isReverse)
					if valid == false {
						break
					}
				default:
					fmt.Println("error!!",com)
					valid = false
					break
				}
			}
		}


		if valid == false {
			fmt.Println("error")
		} else {
			printArray(array, isReverse)
		}
	}
	return
}
