package leetcode

import (
	"AlgorithmWithGo/myLibs"
	"strconv"
)

func PrintTree(root *myLibs.TreeNode) [][]string {
	return printTree(root)
}

func printTree(root *myLibs.TreeNode) [][]string {
	m := make([]interface{}, 0)
	m = append(m, root)
	out := make([]string, 0)
	for {
		switch poped := m[0].(type) {
		case *myLibs.TreeNode:
			//fmt.Println(poped.Val)
			out = append(out, strconv.Itoa(poped.Val))
			if poped.Left != nil {
				m = append(m, poped.Left)
			} else {
				m = append(m, "")
			}
			if poped.Right != nil {
				m = append(m, poped.Right)
			} else {
				m = append(m, "")
			}
		case string:
			//fmt.Println(poped)
			out = append(out, poped)
			m = append(m, "")
			m = append(m, "")
			break
		}
		m = m[1:]
		isBreak := true
		for _, e := range m {
			if _, ok := e.(*myLibs.TreeNode); ok {
				isBreak = false
				break
			}
		}
		if isBreak {
			break
		}
	}
	// fmt.Println(out)

	l := 0
	base := 1
	for {
		if len(out) <= base {
			break
		}
		l += 1
		base += pow(2, l)
	}
	//fmt.Println("total level", l+1)
	width := 2 * pow(2, l) - 1

	str := make([][]string, l+1)
	for i := range str {
		str[i] = make([]string, width)
	}

	level := 0
	x := 0
	for len(out) > 0 {
		d := pow(2, l) - 1
		x = d
		for i := 0 ; i < pow(2, level) ; i++ {
			str[level][x] = getFirst(&out)
			x += 2 * d + 2
		}
		l -= 1
		level += 1
	}

	//for i := range str {
	//	fmt.Println(str[i])
	//}

	return str
}

func getFirst(ss *[]string) string {
	if len(*ss) == 0 {
		return ""
	}
	ans := (*ss)[0]
	*ss = (*ss)[1:]
	return ans
}


func pow(a, b int) int {
	ans := 1
	for i := 0 ; i < b ; i++ {
		ans *= a
	}
	return ans
}