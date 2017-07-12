package leetcode

func TeemoAttacking(timeSeries []int, duration int) int {
	return findPoisonedDuration(timeSeries, duration)
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	result := duration * len(timeSeries)
	diff := 0

	for i := 1 ; i < len(timeSeries) ; i++ {
		diff = timeSeries[i] - timeSeries[i-1]
		if diff < duration {
			result -= duration - diff
		}
	}

	return result
}