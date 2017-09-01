package leetcode

func CombinationSum4(nums []int, target int) int {
	return combinationSum4(nums, target)
}

func combinationSum4(nums []int, target int) int {
	mem := make(map[int]int)
	return comb(nums, target, mem)
}

func comb(nums []int, target int, mem map[int]int) int {
	count := 0

	if target == 0 {
		return 1
	}

	for _, num := range nums {
		if target - num >= 0 {
			temp := 0
			if n, isExist := mem[target-num]; isExist {
				temp = n
			} else {
				mem[target-num] = comb(nums, target - num, mem)
				temp = mem[target-num]
			}
			count += temp
		}
	}

	return count
}