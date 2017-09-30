package leetcode

func FlipLights(n int, m int) int {
	return flipLights(n, m)
}

func flipLights(n int, m int) int {
	switch {
	case m == 0:
		return 1
	case n == 1:
		return 2
	case n == 2 && m == 1:
		return 3
	case n == 2:
		return 4
	case m == 1:
		return 4
	case m == 2:
		return 7
	case m >= 3:
		return 8
	}
	return 8
}
