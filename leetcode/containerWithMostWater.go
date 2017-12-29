package leetcode

// MaxArea is a solution for MaxArea of leetcode.
func MaxArea(height []int) int {
	return maxArea(height)
}

const (
	leftSide = iota
	rightSide
	equalSides
)

func maxArea(height []int) int {
	max, left, right := 0, 0, len(height)-1

	for {
		minH, side := minHeight(height[left], height[right])
		area := (right - left) * minH
		if max < area {
			max = area
		}
		if side == leftSide {
			left++
		} else if side == rightSide {
			right--
		} else {
			_, minSide := minHeight(height[left+1], height[right-1])
			if minSide == leftSide {
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
		return a, leftSide
	} else if a == b {
		return a, equalSides
	}
	return b, rightSide
}
