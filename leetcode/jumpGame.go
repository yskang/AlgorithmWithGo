package leetcode

// CanJump is a solution of the problem "55. Jump Game" in leetcode
func CanJump(nums []int) bool {
	return canJumpGreedy(nums)
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
	jumpTo(0, nums, &isEnd, &visitTable)
	if isEnd {
		return true
	}
	return false
}

func jumpTo(current int, nums []int, isEnd *bool, visitTable *[]bool) {
	if *isEnd == true {
		return
	}
	if (*visitTable)[current] {
		return
	}
	(*visitTable)[current] = true
	if current >= len(nums)-1 {
		*isEnd = true
		return
	}
	for i := 1; i <= nums[current]; i++ {
		jumpTo(current+i, nums, isEnd, visitTable)
	}
}
