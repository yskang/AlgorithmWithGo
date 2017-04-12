package leetcode

func BrickWall(wall [][]int) int {
	return leastBricks(wall)
}

func leastBricks(wall [][]int) int {
	lengthFreqMap := make(map[int]int)
	maxEqualLength := 0
	for _, row := range wall {
		currentLength := 0
		for index, brick := range row[:len(row)-1] {
			row[index] = currentLength + brick
			currentLength = row[index]
			lengthFreqMap[currentLength] += 1
			if lengthFreqMap[currentLength] > maxEqualLength {
				maxEqualLength = lengthFreqMap[currentLength]
			}
		}
	}

	return len(wall) - maxEqualLength
}