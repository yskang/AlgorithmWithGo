// https://www.acmicpc.net/problem/11401

package baekjoon

import (
	"fmt"
)

func binomialCoefficientFermat() {
	var n, r uint64

	fmt.Scanf("%d %d\n", &n, &r)

	p := uint64(1000000007)

	ans := uint64(1)
	t1 := uint64(1)
	t2 := uint64(1)

	for i := uint64(1); i <= n; i++ {
		t1 = t1 * i
		t1 = t1 % p
	}

	for i := uint64(1); i <= r; i++ {
		t2 = t2 * i
		t2 = t2 % p
	}

	for i := uint64(1); i <= n - r; i++ {
		t2 = t2 * i
		t2 = t2 % p
	}

	t3 := mul(t2, p - 2, p)
	t3 = t3 % p
	ans = t1 * t3
	ans = ans % p

	fmt.Println(ans)
}

func mul(x uint64, y uint64, p uint64) uint64 {
	ans := uint64(1)

	for y > 0 {
		if y % 2 != 0 {
			ans = ans * x
			ans = ans % p
		}
		x = x * x
		x = x % p
		y = y / 2
	}
	return ans
}
