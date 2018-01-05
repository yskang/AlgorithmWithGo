package leetcode

// SearchRange is a solution of the problem "34. Search for a Range" in leetcode
func SearchRange(nums []int, target int) []int {
	return searchRange(nums, target)
}

func searchRange(nums []int, target int) []int {
	m := make(map[int][]int)
	for i, n := range nums {
		if n == target {
			m[n] = append(m[n], i)
		} else if n > target {
			break
		}
	}

	if is, isExist := m[target]; isExist {
		return []int{is[0], is[len(is)-1]}
	}

	return []int{-1, -1}
}
