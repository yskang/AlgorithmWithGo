// https://www.acmicpc.net/problem/11051

package baekjoon

import "fmt"

type BC struct{
	N int
	K int
}

var bcMap map[BC]int64

func BinomialCoefficientByDP() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)

	bcMap = make(map[BC]int64)

	bcMap[BC{0, 0}] = 1
	bcMap[BC{1, 1}] = 1
	bcMap[BC{1, 0}] = 1

	fmt.Println(getBc(n, k))
}


func getBc(n int, k int) int64 {
	if n < 0 || k < 0 {
		return 0
	}

	if n == 0 || k == n {
		return 1
	}

	result, ok := bcMap[BC{n, k}]

	if ok {
		return result
	} else {
		bcMap[BC{n, k}] = getBc(n-1, k-1) + getBc(n-1, k)
		return bcMap[BC{n, k}] % 10007
	}
}
