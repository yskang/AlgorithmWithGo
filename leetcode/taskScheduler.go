package leetcode

import (
	"log"
)

func LeastInterval(tasks []byte, n int) int {
	return leastInterval(tasks, n)
}

func leastInterval(tasks []byte, n int) int {
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