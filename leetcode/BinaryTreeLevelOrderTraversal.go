package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return [][]int{[]int{root.Val}}
	}

	return append([][]int{[]int{root.Val}}, myMerge(levelOrder(root.Left), levelOrder(root.Right))...)
}


func myMerge(a [][]int, b [][]int)[][]int {
	lenA := len(a)
	lenB := len(b)

	result := [][]int{}

	if lenA == lenB {
		for i := 0 ; i < lenA ; i++ {
			result = append(result, append(a[i], b[i]...))
		}
	} else if lenA > lenB {
		for i := 0 ; i < lenB ; i++ {
			result = append(result, append(a[i], b[i]...))
		}
		for i := lenB ; i < lenA ; i++ {
			result = append(result, a[i])
		}
	} else {
		for i := 0 ; i < lenA ; i++ {
			result = append(result, append(a[i], b[i]...))
		}
		for i := lenA ; i < lenB ; i++ {
			result = append(result, b[i])
		}
	}
	return result
}
