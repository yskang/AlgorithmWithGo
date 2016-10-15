// https://www.acmicpc.net/problem/1003

package baekjoon

import "fmt"

var count_0, count_1 int

func FibonacciRecursionCount() {
	var n, nn int
	fmt.Scanf("%d/n", &n)

	for i := 0 ; i < n ; i++ {
		fmt.Scanf("%d/n", &nn)
		fibonacci(nn)
		fmt.Println(count_0, count_1)
		count_0, count_1 = 0, 0
	}

}

func fibonacci(n int) int {
	if n == 0 {
		count_0 += 1
		return 0;
	} else if n == 1 {
		count_1 += 1
		return 1;
	} else {
		return fibonacci(n - 1) + fibonacci(n - 2)
	}
}
