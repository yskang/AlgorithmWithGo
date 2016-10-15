// https://www.acmicpc.net/problem/2579

package baekjoon

import "fmt"

func StepUpDynamicProgramming() {
	var numStep, score int
	fmt.Scanf("%d\n", &numStep)

	step := make([]int, 0)
	scores := make([]int, 300)

	for i := 0 ; i < numStep ; i++ {
		fmt.Scanf("%d\n", &score)
		step = append(step, score)
	}

	scores[0] = step[0]
	scores[1] = step[1] + step[0]
	scores[2] = step[2] + maxStep(scores[0], step[1])

	for i := 3 ; i < numStep ; i++ {
		scores[i] = step[i] + maxStep(scores[i-2], scores[i-3] + step[i-1])
	}

	fmt.Println(scores[numStep-1])
}

func maxStep(a int, b int) int {
	if a > b {
		return a
	}
	return b
}