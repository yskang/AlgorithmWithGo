package leetcode

import "fmt"
var cnt int

func solve(n int, a int, b int, c int, chkMap map[string]int) {
	if a + b + c == n {
		str := fmt.Sprintf("%d,%d,%d",a,b,c)
		if a < b + c && b < a + c && c < a + b && a <= b && b <= c && chkMap[str] == 0 {
			cnt += 1
			chkMap[str] = 1
		}
		return
	}

	solve(n, a+1, b, c, chkMap)
	solve(n, a, b+1, c, chkMap)
	solve(n, a, b, c+1, chkMap)
}

func Triangle(n int) int {
	chkMap := make(map[string]int)
	solve(n, 1, 1, 1, chkMap)
	return cnt
}
