package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func LongestCommonSubstring() {
	MAXN := 200020
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	strA := scanner.Text()

	scanner.Scan()
	strB := scanner.Text()

	str := strA + "$" + strB

	group := make([]int, 100)
	tempGroup := make([]int, MAXN)
	team := 1
	n := len(str)
	m := len(strA)
	suffixArray := make([]int, MAXN)

	for i := 0; i < n; i++ {
		suffixArray[i] = i
		group[i] = int(str[i] - 'a')
	}

	comp := func(i, j int) bool {
		x, y := suffixArray[i], suffixArray[j]
		if group[x] == group[y] {
			return group[x+team] < group[y+team]
		}
		return group[x] < group[y]
	}

	for team <= n {
		group[n] = -1
		sort.Slice(suffixArray, comp)
		tempGroup[suffixArray[0]] = 0

		for i := 1; i < n; i++ {
			if comp(i-1, i) {
				tempGroup[suffixArray[i]] = tempGroup[suffixArray[i-1]] + 1
			} else {
				tempGroup[suffixArray[i]] = tempGroup[suffixArray[i-1]]
			}
		}

		for i := 0; i < n; i++ {
			group[i] = tempGroup[i]
		}
		team <<= 1
	}

	rank := make([]int, MAXN)
	LCP := make([]int, MAXN)

	for i := 0; i < n; i++ {
		rank[suffixArray[i]] = i
	}
	length := 0
	for i := 0; i < n; i++ {
		k := rank[i]
		if k != 0 {
			j := suffixArray[k-1]
			for j+length < len(str) && i+length < len(str) && str[j+length] == str[i+length] {
				length++
			}
			LCP[k] = length
			if length != 0 {
				length--
			}
		}
	}

	res := make([]string, MAXN)
	ans := 0
	for i := 1; i < n; i++ {
		if (suffixArray[i-1] > m && suffixArray[i] < m) || (suffixArray[i-1] < m && suffixArray[i] > m) {
			if ans < LCP[i] {
				ans = LCP[i]
				for j := 0; j < ans; j++ {
					res[j] = string(str[j+suffixArray[i]])
				}
			}
		}
	}
	fmt.Println(ans)
	fmt.Println(strings.Join(res[:ans+1], ""))
}
