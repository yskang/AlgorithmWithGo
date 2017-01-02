package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	sequence := "1"
	tempSequence := ""

	for i := 1 ; i < n ; i++ {
		current := sequence[0]
		count := 1
		for j := 1 ; j < len(sequence) ; j++ {
			if current != sequence[j] {
				tempSequence += strconv.Itoa(count) + string(current)
				current = sequence[j]
				count = 1
			} else {
				count += 1
			}
		}
		tempSequence += strconv.Itoa(count) + string(current)
		sequence = tempSequence
		tempSequence = ""
	}

	return sequence
}
