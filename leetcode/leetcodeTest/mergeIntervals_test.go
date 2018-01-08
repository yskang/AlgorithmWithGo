package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleMergeIntervals() {
	intervals := make([]leetcode.Interval, 0)
	intervals = append(intervals, leetcode.Interval{1, 3})
	intervals = append(intervals, leetcode.Interval{2, 6})
	intervals = append(intervals, leetcode.Interval{8, 10})
	intervals = append(intervals, leetcode.Interval{15, 18})
	fmt.Println(leetcode.MergeIntervals(intervals))
	intervals = []leetcode.Interval{}
	intervals = append(intervals, leetcode.Interval{1, 4})
	intervals = append(intervals, leetcode.Interval{4, 5})
	fmt.Println(leetcode.MergeIntervals(intervals))
	intervals = []leetcode.Interval{}
	intervals = append(intervals, leetcode.Interval{1, 4})
	intervals = append(intervals, leetcode.Interval{0, 4})
	fmt.Println(leetcode.MergeIntervals(intervals))
	intervals = []leetcode.Interval{}
	intervals = append(intervals, leetcode.Interval{4, 5})
	intervals = append(intervals, leetcode.Interval{2, 4})
	intervals = append(intervals, leetcode.Interval{4, 6})
	intervals = append(intervals, leetcode.Interval{3, 4})
	intervals = append(intervals, leetcode.Interval{0, 0})
	intervals = append(intervals, leetcode.Interval{1, 1})
	intervals = append(intervals, leetcode.Interval{3, 5})
	intervals = append(intervals, leetcode.Interval{2, 2})
	fmt.Println(leetcode.MergeIntervals(intervals))
	// output:
	// [{1 6} {8 10} {15 18}]
	// [{1 5}]
	// [{0 4}]
	// [{0 0} {1 1} {2 6}]
}
