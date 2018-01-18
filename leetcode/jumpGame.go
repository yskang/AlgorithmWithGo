package leetcode

// CanJump is a solution of the problem "55. Jump Game" in leetcode
func CanJump(nums []int) bool {
	return canJump(nums)
}

func canJumpGreedy(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

func canJump(nums []int) bool {
	isEnd := false
	visitTable := make([]bool, len(nums))
	if len(nums) == 1 {
		return true
	}
	jumpTo(0, nums, &isEnd, &visitTable)
	return isEnd
}

func jumpTo(current int, nums []int, isEnd *bool, visitTable *[]bool) {
	(*visitTable)[current] = true
	for i := nums[current]; i >= 1; i-- {
		if current+i >= len(nums)-1 {
			*isEnd = true
			return
		}
		if !(*visitTable)[current+i] {
			jumpTo(current+i, nums, isEnd, visitTable)
		}
	}
}
