package leetcode

func MaxArea(height []int) int {
	return maxArea(height)
}

const LEFT = 0
const RIGHT = 1
const EQ = 2

func maxArea(height []int) int {
	max, left, right := 0, 0, len(height)-1

	for {
		minH, side := minHeight(height[left], height[right])
		area := (right - left) * minH
		if max < area {
			max = area
		}
		if side == LEFT {
			left++
		} else if side == RIGHT {
			right--
		} else {
			_, minSide := minHeight(height[left+1], height[right-1])
			if minSide == LEFT {
				right--
			} else {
				left++
			}
		}

		if left == right {
			break
		}
	}
	return max
}

func minHeight(a, b int) (int, int) {
	if a < b {
		return a, LEFT
	} else if a == b {
		return a, EQ
	}
	return b, RIGHT
}
