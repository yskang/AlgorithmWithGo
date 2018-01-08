package leetcodeTest

import (
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func ExampleMergeIntervals() {
	intervals := make([]leetcode.Interval, 0)
	intervals = append(intervals, leetcode.Interval{Start: 1, End: 3})
	intervals = append(intervals, leetcode.Interval{Start: 2, End: 6})
	intervals = append(intervals, leetcode.Interval{Start: 8, End: 10})
	intervals = append(intervals, leetcode.Interval{Start: 15, End: 18})
	fmt.Println(leetcode.MergeIntervals(intervals))
	intervals = []leetcode.Interval{}
	intervals = append(intervals, leetcode.Interval{Start: 1, End: 4})
	intervals = append(intervals, leetcode.Interval{Start: 4, End: 5})
	fmt.Println(leetcode.MergeIntervals(intervals))
	intervals = []leetcode.Interval{}
	intervals = append(intervals, leetcode.Interval{Start: 1, End: 4})
	intervals = append(intervals, leetcode.Interval{Start: 0, End: 4})
	fmt.Println(leetcode.MergeIntervals(intervals))
	intervals = []leetcode.Interval{}
	intervals = append(intervals, leetcode.Interval{Start: 4, End: 5})
	intervals = append(intervals, leetcode.Interval{Start: 2, End: 4})
	intervals = append(intervals, leetcode.Interval{Start: 4, End: 6})
	intervals = append(intervals, leetcode.Interval{Start: 3, End: 4})
	intervals = append(intervals, leetcode.Interval{Start: 0, End: 0})
	intervals = append(intervals, leetcode.Interval{Start: 1, End: 1})
	intervals = append(intervals, leetcode.Interval{Start: 3, End: 5})
	intervals = append(intervals, leetcode.Interval{Start: 2, End: 2})
	fmt.Println(leetcode.MergeIntervals(intervals))
	// output:
	// [{1 6} {8 10} {15 18}]
	// [{1 5}]
	// [{0 4}]
	// [{0 0} {1 1} {2 6}]
}
