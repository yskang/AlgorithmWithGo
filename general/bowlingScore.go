package general

import (
	"strconv"
	"errors"
)

func BowlingScore(pins string) int {
	score := 0
	isFirst := true
	firstScore := 0

	for i := 0 ; i < len(pins) ; i++ {
		if pins[i] == '/' || pins[i] == 'X' {
			bonusLength := 2
			if pins[i] == '/' {
				bonusLength = 1
			}
			bonus, err := getBonusScore(pins, i + 1, bonusLength)
			if err != nil {
				return score
			}
			score += 10 + bonus
			isFirst = true
		} else {
			pin, _ := strconv.Atoi(string(pins[i]))

			if isFirst {
				isFirst = false
				firstScore = pin
			} else {
				isFirst = true
				score += firstScore + pin
			}
		}
	}
	return score
}

func getBonusScore(pins string, start int, length int) (int, error) {
	if start + length > len(pins) {
		return 0, errors.New("open")
	}

	bonus := 0

	for i := 0 ; i < length ; i++ {
		pin, err := strconv.Atoi(string(pins[start + i]))
		if err != nil{
			if pins[start + i] == 'X' {
				pin = 10
			} else if pins[start + i] == '/' {
				return 10, nil
			}
		}
		bonus += pin
	}

	return bonus, nil
}
