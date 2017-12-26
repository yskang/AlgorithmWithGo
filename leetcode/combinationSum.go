package leetcode

// CombinationSum get combincation sum, leetcode 39
func CombinationSum(candidates []int, target int) [][]int {
	return combinationSum(candidates, target)
}

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	getCombiSum(candidates, target, &ans, []int{}, 0)
	return ans
}

func getCombiSum(candidates []int, target int, ans *[][]int, single []int, startIndex int) {
	if target == 0 {
		// log.Println("add", single, "to ans", *ans)
		*ans = append(*ans, append([]int{}, single...))
	} else if target > 0 {
		for i := startIndex; i < len(candidates); i++ {
			// log.Println("add", candidates[i], "to single", single)
			getCombiSum(candidates, target-candidates[i], ans, append(single, candidates[i]), i)
		}
	}
}
