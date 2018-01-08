package leetcode

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

//MergeIntervals is a solution of the problem "56. Merge Intervals" in leetcode
func MergeIntervals(intervals []Interval) []Interval {
	return mergeIntervals(intervals)
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start < intervals[j].Start {
			return true
		}
		return false
	})

	merged := make([]Interval, 0)

	merged = append(merged, intervals[0])
	for _, interval := range intervals[1:] {
		lastInterval := merged[len(merged)-1]
		if lastInterval.End < interval.Start {
			merged = append(merged, interval)
		} else {
			maxLast := lastInterval.End
			if lastInterval.End < interval.End {
				maxLast = interval.End
			}
			merged[len(merged)-1] = Interval{lastInterval.Start, maxLast}
		}
	}

	return merged
}

func mergeIntervals2(intervals []Interval) []Interval {
	newInterval := make([]Interval, 0)
	type point struct {
		position int
		state    bool
	}
	points := make([]point, 0)
	for _, interval := range intervals {
		points = append(points, point{interval.Start, true})
		points = append(points, point{interval.End, false})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].position < points[j].position {
			return true
		} else if points[i].position == points[j].position {
			if points[i].state == true && points[j].state == false {
				return true
			}
		}
		return false
	})

	start := 0
	startPosition := 0
	for _, p := range points {
		if p.state == true {
			if start == 0 {
				startPosition = p.position
			}
			start++
		} else {
			start--
		}
		if start == 0 {
			var lastInterval Interval
			if len(newInterval) > 0 {
				lastInterval = newInterval[len(newInterval)-1]
				if lastInterval.End == startPosition {
					newInterval[len(newInterval)-1] = Interval{lastInterval.Start, p.position}
				} else {
					newInterval = append(newInterval, Interval{startPosition, p.position})
				}
			} else {
				newInterval = append(newInterval, Interval{startPosition, p.position})
			}
		}
	}

	return newInterval
}
