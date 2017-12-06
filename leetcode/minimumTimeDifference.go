package leetcode

import (
	"sort"
	"strings"
	"strconv"
)

func FindMinDifference(timePoints []string) int {
	return findMinDifference(timePoints)
}

func findMinDifference(timePoints []string) int {
	sort.Slice(timePoints, func(i, j int) bool {
		return isSmall(timePoints[i], timePoints[j])
	})

	minDiff := getDiff("00:00", timePoints[0]) + getDiff(timePoints[len(timePoints)-1], "24:00")
	prevT := timePoints[0]
	for _, t := range timePoints[1:] {
		diff := getDiff(prevT, t)
		if diff < minDiff {
			minDiff = diff
		}
		prevT = t
	}

	return minDiff
}

func getDiff(t1 string, t2 string) int {
	hour1, min1 := getTime(t1)
	hour2, min2 := getTime(t2)
	minDiff := min2 - min1
	hourDiff := hour2 - hour1
	return hourDiff * 60 + minDiff
}

func isSmall(a string, b string) bool {
	hourA, minA := getTime(a)
	hourB, minB := getTime(b)

	if hourA < hourB {
		return true
	} else if hourA > hourB {
		return false
	}

	return minA < minB
}

func getTime(timeStr string) (int, int) {
	timeStrs := strings.Split(timeStr, ":")
	h, _ := strconv.ParseInt(timeStrs[0], 10, 64)
	m, _ := strconv.ParseInt(timeStrs[1], 10, 64)
	return int(h), int(m)
}