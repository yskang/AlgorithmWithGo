package leetcode

import (
	"log"
	"math"
	"sort"
)

func LeastInterval(tasks []byte, n int) int {
	return leastInterval(tasks, n)
}

func leastInterval(tasks []byte, n int) int {
	list := make([]int, 26)
	for _, c := range tasks {
		list[c-'A'] += 1
	}
	sort.Ints(list)
	maxVal := list[25]-1
	idleSlots := maxVal * n
	for i := 24 ; i >= 0 && list[i] > 0 ; i-- {
		if list[i] < maxVal {
			idleSlots -= list[i]
		} else {
			idleSlots -= maxVal
		}
	}
	if idleSlots > 0 {
		return idleSlots + len(tasks)
	}
	return len(tasks)
}

func leastInterval_(tasks []byte, n int) int {
	taskMap := make(map[byte]int)
	for _, task := range tasks {
		taskMap[task] += 1
	}
	maxTaskCount := 0
	for _, count := range taskMap {
		if count > maxTaskCount {
			maxTaskCount = count
		}
	}
	restTaskCount := len(tasks) - maxTaskCount

	log.Println("max count:", maxTaskCount)
	log.Println("rest count:", restTaskCount)

	if restTaskCount == n * maxTaskCount {
		log.Println("1")
		return len(tasks)
	} else if restTaskCount < n * maxTaskCount {
		log.Println("2")
		return maxTaskCount * (n + 1) - (n - int(math.Floor(float64(restTaskCount/maxTaskCount))))
	} else if restTaskCount > n * maxTaskCount {
		if restTaskCount % maxTaskCount == 0 {
			log.Println("3")
			return len(tasks)
		} else {
			log.Println("4")
			return (int(math.Floor(float64(restTaskCount / maxTaskCount))) + 2) * maxTaskCount
		}
	}
	return 0
}

func leastInterval_slow(tasks []byte, n int) int {
	taskMap := make(map[byte]int)
	coolMap := make(map[byte]int)

	for _, taskName := range tasks {
		taskMap[taskName] += 1
		coolMap[taskName] = 0
	}

	log.Println("taskMap", taskMap)
	log.Println("coolMap", coolMap)

	count := 0
	for {
		task, isEnd := getMaximumAvailableTask(&taskMap, &coolMap, n)
		log.Println(string(task))
		if isEnd {
			log.Println("result", count)
			return count
		}
		count += 1
	}

	return 0
}

func getMaximumAvailableTask(tasks *map[byte]int, coolTime *map[byte]int, n int) (byte, bool) {
	maxCount := 0
	maxTasks := make([]byte, 0)

	log.Println("tasks", *tasks)

	for task, count := range *tasks {
		if count > maxCount && (*coolTime)[task] <= 0 {
			maxCount = count
			maxTasks = append([]byte{}, task)
		} else if count > 0 && count == maxCount && (*coolTime)[task] <= 0 {
			maxTasks = append(maxTasks, task)
		}
	}

	log.Println("maxTasks", maxTasks)

	if len(maxTasks) == 0 {
		for task, count := range *tasks {
			if count > maxCount {
				maxCount = count
				maxTasks = append([]byte{}, task)
			} else if count > 0 && count == maxCount {
				maxTasks = append(maxTasks, task)
			}
		}
		if len(maxTasks) == 0 {
			return 'n', true
		}
	}

	for _, task := range maxTasks {
		if (*coolTime)[task] <= 0 {
			for taskName, _ := range *coolTime {
				(*coolTime)[taskName] -= 1
			}
			(*coolTime)[task] = n
			(*tasks)[task] -= 1
			return task, false
		}
	}

	for taskName, _ := range *coolTime {
		(*coolTime)[taskName] -= 1
	}

	return 'i', false

}