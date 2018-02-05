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

	sa := make([]int, len(str))
	bucket := make([]int, len(str))

	for i := range str {
		sa[i] = i
		bucket[i] = i
	}

	sa = sortBucket(str, bucket, 1)

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
		prev = line
	}

	fmt.Println(maxLen)
	fmt.Println(str[maxStart : maxStart+maxLen])
}

func sortBucket(target string, bucket []int, order int) []int {
	d := make(map[string][]int)
	dkey := []string{}
	for _, i := range bucket {
		to := i + order
		if to >= len(target) {
			to = len(target)
		}
		key := target[i:to]
		d[key] = append(d[key], i)
	}
	result := []int{}

	for k := range d {
		dkey = append(dkey, k)
	}

	sort.Strings(dkey)
	for _, k := range dkey {
		if len(d[k]) > 1 {
			result = append(result, sortBucket(target, d[k], order*2)...)
		} else {
			result = append(result, d[k][0])
		}
	}
	return result
}
