package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func LongestCommonSubstring() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	strA := scanner.Text()
	scanner.Scan()
	strB := scanner.Text()

	str := strA + "!" + strB

	n := len(str)
	lenA := len(strA)
	// lenB := len(strB)

	sa := make([]int, len(str))
	w := 1

	for i := range str {
		sa[i] = i
	}

	for w <= n {
		sort.Slice(sa, func(i, j int) bool {
			start, end := sa[i], sa[i]+w
			if sa[i]+w > n {
				end = n
			}
			a := str[start:end]
			start = sa[j]
			end = sa[j] + w
			if sa[j]+w > n {
				end = n
			}
			b := str[start:end]
			return a < b
		})
		w *= 2
	}

	maxLen := 0
	maxStart := 0
	minLen := 0
	prev := sa[0]
	prevSide, currSide := false, false //front false, back true

	for _, line := range sa[1:] {
		if prev < lenA {
			prevSide = false
		} else {
			prevSide = true
		}
		if line < lenA {
			currSide = false
		} else {
			currSide = true
		}
		if prevSide != currSide {
			if n-prev < n-line {
				minLen = n - prev
			} else {
				minLen = n - line
			}
			for i := 0; i < minLen; i++ {
				if str[prev+i] != str[line+i] {
					if i > maxLen {
						maxLen = i
						maxStart = prev
					}
					break
				}
				if i == minLen-1 {
					if i+1 > maxLen {
						maxLen = i + 1
						maxStart = line
					}
				}
			}
		}
		// fmt.Println(str[prev:], str[line:], prevSide, currSide, maxLen)
		prev = line
	}

	fmt.Println(maxLen)
	fmt.Println(str[maxStart : maxStart+maxLen])
}
