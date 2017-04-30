package beakjoon

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
	"bufio"
	"os"
)

func ATM() {
	var numberOfPerson int
	fmt.Scanf("%d\n", &numberOfPerson)
	reader := bufio.NewReader(os.Stdin)
	waitingString, _ := reader.ReadString('\n')
	waitingString = strings.TrimRight(waitingString, "\n")
	waitingTimes := make([]int, numberOfPerson)
	waits := strings.Split(waitingString, " ")
	for i, wait := range waits {
		num, err := strconv.Atoi(wait)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			waitingTimes[i] = num
		}
	}
	sort.Ints(waitingTimes)

	sum := 0
	for i, waitTime := range waitingTimes {
		sum += (numberOfPerson - i) * waitTime
	}

	fmt.Println(sum)
}
