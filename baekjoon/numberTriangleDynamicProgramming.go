// https://www.acmicpc.net/problem/1932

package baekjoon

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func NumberTriangleDynamicPrograming() {
	triangle := make([][]int, 0)
	var n int
	fmt.Scanf("%d\n", &n)

	//for i := 0 ; i < n ; i++ {
	//	triangle = append(triangle, []int{})
	//	for j := 0 ; j <= i ; j++ {
	//		fmt.Scanf("%d", &m)
	//		triangle[i] = append(triangle[i], m)
	//	}
	//}

	inNums := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	i := 0

	for scanner.Scan() {
		inline := scanner.Text()
		inStrs := strings.Split(inline, " ")

		triangle = append(triangle, []int{})
		for _, str := range inStrs {
			num, _ := strconv.Atoi(str)
			triangle[i] = append(triangle[i], num)
		}

		inNums = inNums[:0]
		i += 1

		if i == n {
			break
		}
	}

	for i = n-2 ; i >= 0 ; i-- {
		for j := 0 ; j < i+1 ; j++ {
			triangle[i][j] = triangle[i][j] + max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	fmt.Println(triangle[0][0])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

