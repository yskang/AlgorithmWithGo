// https://algospot.com/judge/problem/read/SKICOURSE

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func main() {


	var T int
	fmt.Scanf("%d\n", &T)

	scanner := bufio.NewScanner(os.Stdin)

	for t := 0 ; t < T ; t++ {

		var inStr string
		for scanner.Scan() {
			inStr = scanner.Text()
			break
		}
		inLine := strings.Split(inStr, " ")
		var N uint16
		var M uint32
		var S uint8
		var temp int

		temp, _ = strconv.Atoi(inLine[0])
		N = uint16(temp)
		temp, _ = strconv.Atoi(inLine[1])
		M = uint32(temp)
		temp, _ = strconv.Atoi(inLine[2])
		S = uint8(temp)

		graph := make([]map[uint16]int32, 0)
		for i := uint16(0) ; i < N ; i++ {
			graph = append(graph, make(map[uint16]int32))
		}

		cache := make([][]int32, 0)
		for i := uint16(0) ; i < N ; i++ {
			cache = append(cache, make([]int32, S))
		}

		for m := uint32(0) ; m < M ; m++ {
			var inStr2 string
			for scanner.Scan() {
				inStr2 = scanner.Text()
				break
			}
			inLine2 := strings.Split(inStr2, " ")
			A, _ := strconv.Atoi(inLine2[0])
			B, _ := strconv.Atoi(inLine2[1])
			C, _ := strconv.Atoi(inLine2[2])
			graph[B-1][uint16(A-1)] = int32(C)
		}

		maxExcitingValue := int32(math.MinInt32)
		var excitingValue int32

		for slope := uint8(1) ; slope <= S ; slope++ {
			for node := uint16(0) ; node < N ; node++ {

				excitingValue = getMaxExcitingValueAt(&cache, &graph, node, slope)

				cache[node][slope-1] = excitingValue

				if maxExcitingValue < excitingValue {
					maxExcitingValue = excitingValue
				}

			}
		}

		fmt.Println(maxExcitingValue)
	}
}

func getMaxExcitingValueAt(cache *[][]int32, graph *[]map[uint16]int32, node uint16, depth uint8) int32 {
	maxExcitingValue := int32(math.MinInt32)
	excitingValue := int32(0)

	if len((*graph)[node]) == 0 {
		return 0
	}

	for prevNode := range (*graph)[node] {
		if depth == 1 {
			excitingValue = (*graph)[node][prevNode]
		} else {
			if (*cache)[prevNode][depth-2] != 0 {
				excitingValue = (*cache)[prevNode][depth-2] + (*graph)[node][prevNode]
			} else {
				excitingValue = getMaxExcitingValueAt(cache, graph, prevNode, depth-1) + (*graph)[node][prevNode]
			}
		}
		if maxExcitingValue < excitingValue {
			maxExcitingValue = excitingValue
		}
	}

	(*cache)[node][depth-1] = maxExcitingValue
	return maxExcitingValue
}